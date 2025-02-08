package data_provider

import (
	"context"

	"likemind/internal/db/model"
	"likemind/internal/db/op"
)

type DataProvider[M model.M[F, PK], F model.F, PK comparable] interface {
	Get(ctx context.Context) GetBuilder[M, F, PK]
	List(ctx context.Context) ListBuilder[M, F, PK]
	Insert(ctx context.Context) InsertBuilder[M, F, PK]
	Delete(ctx context.Context) DeleteBuilder[M, F, PK]
	Update(ctx context.Context) UpdateBuilder[M, F, PK]
}

type ListBuilder[M model.M[F, PK], F model.F, PK comparable] interface {
	Filter(field F, op op.Operator, val any) ListBuilder[M, F, PK]
	Do(ctx context.Context) ([]M, error)
}

type GetBuilder[M model.M[F, PK], F model.F, PK comparable] interface {
	ByPK(pk PK) GetBuilder[M, F, PK]
	ByFilter(field F, op op.Operator, val any) GetBuilder[M, F, PK]
	Do(ctx context.Context) (M, error)
}

type UpdateBuilder[M model.M[F, PK], F model.F, PK comparable] interface {
	Set(field F, val any) UpdateBuilder[M, F, PK]
	WherePK(pk PK) UpdateBuilder[M, F, PK]
	Do(ctx context.Context) error
}

type InsertBuilder[M model.M[F, PK], F model.F, PK comparable] interface {
	Field(field F, val any) InsertBuilder[M, F, PK]
	Do(ctx context.Context) (PK, error)
}

type DeleteBuilder[M model.M[F, PK], F model.F, PK comparable] interface {
	ByPK(pk PK) DeleteBuilder[M, F, PK]
	ByFilter(field F, op op.Operator, val any) DeleteBuilder[M, F, PK]
	Do(ctx context.Context) error
}
