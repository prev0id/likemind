package domain

import (
	"context"
	"time"
)

type Data interface {
	GetID() int64
}

type DataProvider[T Data] interface {
	Get(ctx context.Context, id int64) (T, error)
	List(ctx context.Context) ([]T, error)
	Insert(ctx context.Context, data T) (int64, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, data T) error
	UpdateField(ctx context.Context, id int64, field string, value any) error
}

type DataFields struct {
	ID        int64     `db:"id" goqu:"skipupdate,skipinsert"`
	CreatedAt time.Time `db:"created_at" goqu:"skipupdate,skipinsert"`
	UpdatedAt time.Time `db:"updated_at" goqu:"defaultifempty"`
}

func (df DataFields) GetID() int64 {
	return df.ID
}
