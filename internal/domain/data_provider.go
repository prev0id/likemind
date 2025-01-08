package domain

import (
	"context"
	"time"
)

// TODO: fix updated at column

type Data interface {
	GetID() int64
}

type DataProvider[T Data] interface {
	Get(ctx context.Context, field string, value any) (T, error)
	List(ctx context.Context) ([]T, error)
	Insert(ctx context.Context, data T) (int64, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, data T) error
	UpdateField(ctx context.Context, id int64, field string, value any) error
}

type DataFields struct {
	ID        int64     `db:"id" fieldopt:"omitempty"`
	CreatedAt time.Time `db:"created_at" fieldopt:"omitempty"`
	UpdatedAt time.Time `db:"updated_at" fieldopt:"omitempty"`
}

func (df DataFields) GetID() int64 {
	return df.ID
}
