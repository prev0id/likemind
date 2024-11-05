package profile

import (
	"context"

	"likemind/internal/domain"
)

type Service interface {
	UpdateName(ctx context.Context, id uint64)
}

type db interface {
	ListUsers(ctx context.Context) ([]domain.User, error)
	GetUser(ctx context.Context) (domain.User, error)
	UpdateUser(ctx context.Context, user domain.User) error
}

type implementation struct {
	db db
}

func New(db db) Service {
}
