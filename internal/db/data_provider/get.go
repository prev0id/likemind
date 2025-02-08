package data_provider

import (
	"context"
	"strings"

	"likemind/internal/db/model"
	"likemind/internal/db/op"

	"github.com/georgysavva/scany/v2/pgxscan"
)

type pgGetBuilder[M model.M[F, PK], F model.F, PK comparable] struct {
	provider *pgDataProvider[M, F, PK]
	filters  []condition
}

func (b *pgGetBuilder[M, F, PK]) ByPK(pk PK) GetBuilder[M, F, PK] {
	b.filters = append(b.filters, condition{field: b.provider.pkField.String(), op: op.Eq, value: pk})
	return b
}

func (b *pgGetBuilder[M, F, PK]) ByFilter(field F, operator op.Operator, value any) GetBuilder[M, F, PK] {
	b.filters = append(b.filters, condition{field: field.String(), op: operator, value: value})
	return b
}

func (b *pgGetBuilder[M, F, PK]) Do(ctx context.Context) (M, error) {
	query := &strings.Builder{}

	query.WriteString("SELECT * FROM ")
	query.WriteString(b.provider.table.String())
	query.WriteRune(' ')
	args := writeWhere(query, b.filters, argsIndexingStart)

	var result M

	err := pgxscan.Get(ctx, b.provider.pool, &result, query.String(), args...)

	return result, err
}
