package profile

import (
	"context"

	"likemind/internal/database/adapter/profile_adapter"
	"likemind/internal/domain"
)

type Service interface {
	CreateUser(ctx context.Context, user domain.User) (int64, error)
	UpdateUser(ctx context.Context, user domain.User) error
	DeleteProfile(ctx context.Context, id int64) error
	GetProfile(ctx context.Context, id int64) (domain.Profile, error)
}

type implementation struct {
	db profile_adapter.Adapter
}

func New(db profile_adapter.Adapter) Service {
	return &implementation{
		db: db,
	}
}

func (s *implementation) CreateUser(ctx context.Context, user domain.User) (int64, error) {
	return s.db.CreateUser(ctx, user)
}

func (s *implementation) UpdateUser(ctx context.Context, user domain.User) error {
	return s.db.UpdateUser(ctx, user)
}

func (s *implementation) DeleteProfile(ctx context.Context, id int64) error {
	return s.db.RemoveUser(ctx, id)
}

func (s *implementation) GetProfile(ctx context.Context, id int64) (domain.Profile, error) {
	return s.db.GetProfileByUserID(ctx, id)
}
