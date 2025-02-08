package data_provider

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"likemind/internal/db/model"
	"likemind/internal/db/op"
)

type pgUpdateBuilder[M model.M[F, PK], F model.F, PK comparable] struct {
	provider  *pgDataProvider[M, F, PK]
	sets      []setField
	filters   []condition
	tableName string
}

func (b *pgUpdateBuilder[M, F, PK]) WherePK(pk PK) UpdateBuilder[M, F, PK] {
	b.filters = append(b.filters, condition{field: b.provider.pkField.String(), op: op.Eq, value: pk})
	return b
}

func (b *pgUpdateBuilder[M, F, PK]) Set(field F, val any) UpdateBuilder[M, F, PK] {
	b.sets = append(b.sets, setField{field: field.String(), value: val})
	return b
}

func (b *pgUpdateBuilder[M, F, PK]) Do(ctx context.Context) error {
	if len(b.sets) == 0 {
		return errors.New("no fields to update")
	}

	if len(b.filters) == 0 {
		return errors.New("update requires a filter condition")
	}

	b.sets = append(b.sets, setField{field: b.provider.updatedAtField.String(), value: time.Now()})

	query := &strings.Builder{}

	args := make([]any, 0, len(b.filters)+len(b.sets))
	argsShift := argsIndexingStart

	setArgs := writeSets(query, b.sets, argsShift)
	argsShift += len(setArgs)
	args = append(args, setArgs...)

	query.WriteString(" SET ")
	query.WriteString(b.tableName)

	query.WriteRune(' ')

	whereArgs := writeWhere(query, b.filters, argsShift)
	args = append(args, whereArgs...)

	_, err := b.provider.pool.Exec(ctx, query.String(), args...)
	return err
}

type setField struct {
	field string
	value any
}

func writeSets(query *strings.Builder, sets []setField, argShift int) []any {
	args := make([]any, 0, len(sets))

	for i, set := range sets {
		if i != 0 {
			query.WriteString(", ")
		}
		query.WriteString(set.field)
		query.WriteString(" = $")
		query.WriteString(strconv.Itoa(argShift + i))
	}

	return args
}
