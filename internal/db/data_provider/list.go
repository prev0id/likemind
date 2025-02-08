package data_provider

import (
	"context"
	"strings"

	"likemind/internal/db/model"
	"likemind/internal/db/op"

	"github.com/georgysavva/scany/v2/pgxscan"
)

type pgListBuilder[M model.M[F, PK], F model.F, PK comparable] struct {
	provider  *pgDataProvider[M, F, PK]
	filters   []condition
	tableName string
}

func (b *pgListBuilder[M, F, PK]) Filter(field F, operator op.Operator, val any) ListBuilder[M, F, PK] {
	b.filters = append(b.filters, condition{field: field.String(), op: operator, value: val})
	return b
}

func (b *pgListBuilder[M, F, PK]) Do(ctx context.Context) ([]M, error) {
	query := &strings.Builder{}
	query.WriteString("SELECT * FROM ")
	query.WriteString(b.tableName)
	args := writeWhere(query, b.filters, argsIndexingStart)

	var results []M
	err := pgxscan.Select(ctx, b.provider.pool, &results, query.String(), args...)
	return results, err
}
