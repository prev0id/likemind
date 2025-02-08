package data_provider

import (
	"context"
	"errors"
	"strings"

	"likemind/internal/db/model"
	"likemind/internal/db/op"
)

type pgDeleteBuilder[M model.M[F, PK], F model.F, PK comparable] struct {
	provider *pgDataProvider[M, F, PK]
	filters  []condition
}

func (b *pgDeleteBuilder[M, F, PK]) ByPK(pk PK) DeleteBuilder[M, F, PK] {
	b.filters = append(b.filters, condition{field: b.provider.pkField.String(), op: op.Eq, value: pk})
	return b
}

func (b *pgDeleteBuilder[M, F, PK]) ByFilter(field F, operator op.Operator, val any) DeleteBuilder[M, F, PK] {
	b.filters = append(b.filters, condition{field: field.String(), op: operator, value: val})
	return b
}

func (b *pgDeleteBuilder[M, F, PK]) Do(ctx context.Context) error {
	if len(b.filters) == 0 {
		return errors.New("delete requires at least one condition")
	}

	query := &strings.Builder{}
	query.WriteString("DELETE FROM ")
	query.WriteString(b.provider.table.String())
	query.WriteRune(' ')
	args := writeWhere(query, b.filters, argsIndexingStart)

	_, err := b.provider.pool.Exec(ctx, query.String(), args...)
	return err
}
