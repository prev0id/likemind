package data_provider

import (
	"context"

	"likemind/internal/db/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

const argsIndexingStart = 1

type pgDataProvider[M model.M[F, PK], F model.F, PK comparable] struct {
	pool  *pgxpool.Pool
	table model.Table

	pkField        F
	updatedAtField F
	createdAtField F
}

func NewPGDataProvider[M model.M[F, PK], F model.F, PK comparable](pool *pgxpool.Pool) DataProvider[M, F, PK] {
	var m M

	return &pgDataProvider[M, F, PK]{
		pool:           pool,
		table:          m.TableName(),
		pkField:        m.FieldPrimaryKey(),
		createdAtField: m.FieldPrimaryKey(),
		updatedAtField: m.FieldPrimaryKey(),
	}
}

func (p *pgDataProvider[M, F, PK]) Get(ctx context.Context) GetBuilder[M, F, PK] {
	return &pgGetBuilder[M, F, PK]{provider: p}
}

func (p *pgDataProvider[M, F, PK]) List(ctx context.Context) ListBuilder[M, F, PK] {
	return &pgListBuilder[M, F, PK]{provider: p}
}

func (p *pgDataProvider[M, F, PK]) Insert(ctx context.Context) InsertBuilder[M, F, PK] {
	return &pgInsertBuilder[M, F, PK]{provider: p}
}

func (p *pgDataProvider[M, F, PK]) Delete(ctx context.Context) DeleteBuilder[M, F, PK] {
	return &pgDeleteBuilder[M, F, PK]{provider: p}
}

func (p *pgDataProvider[M, F, PK]) Update(ctx context.Context) UpdateBuilder[M, F, PK] {
	return &pgUpdateBuilder[M, F, PK]{provider: p}
}
