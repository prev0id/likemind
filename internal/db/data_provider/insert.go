package data_provider

import (
	"context"
	"fmt"
	"strings"
	"time"

	"likemind/internal/db/model"

	"github.com/georgysavva/scany/v2/pgxscan"
)

type pgInsertBuilder[M model.M[F, PK], F model.F, PK comparable] struct {
	provider *pgDataProvider[M, F, PK]
	fields   []insertValue
}

func (b *pgInsertBuilder[M, F, PK]) Field(field F, value any) InsertBuilder[M, F, PK] {
	b.fields = append(b.fields, insertValue{field: field.String(), value: value})
	return b
}

func (b *pgInsertBuilder[M, F, PK]) Do(ctx context.Context) (PK, error) {
	b.fields = append(
		b.fields,
		insertValue{field: b.provider.updatedAtField.String(), value: time.Now()},
		insertValue{field: b.provider.createdAtField.String(), value: time.Now()},
	)

	query := &strings.Builder{}

	query.WriteString("INSERT INTO ")
	query.WriteString(b.provider.table.String())
	query.WriteRune(' ')

	args := writeValues(query, b.fields, argsIndexingStart)

	query.WriteString(" RETURNING ")
	query.WriteString(b.provider.pkField.String())

	var pk PK
	err := pgxscan.Get(ctx, b.provider.pool, &pk, query.String(), args...)
	return pk, err
}

type insertValue struct {
	field string
	value any
}

func writeValues(query *strings.Builder, values []insertValue, argsShift int) []any {
	columns := make([]string, 0, len(values))
	placeholders := make([]string, 0, len(values))
	args := make([]any, 0, len(values))

	for i, value := range values {
		columns = append(columns, value.field)
		placeholders = append(placeholders, fmt.Sprintf("$%d", argsShift+i))
		args = append(args, value.value)
	}
	query.WriteString("( ")
	query.WriteString(strings.Join(columns, ", "))
	query.WriteString(") VALUES (")
	query.WriteString(strings.Join(placeholders, ", "))
	query.WriteString(")")

	return args
}
