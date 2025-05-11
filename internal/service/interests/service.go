package interests

import (
	"context"

	"likemind/internal/database/adapter/interest_adapter"
	"likemind/internal/domain"
)

type Service interface {
	GetUserInterests(ctx context.Context, id domain.UserID) (domain.Interests, error)
	GetGroupInterests(ctx context.Context, id domain.GroupID) (domain.Interests, error)
}

type implementation struct {
	db interest_adapter.Adapter
}

func New(db interest_adapter.Adapter) Service {
	return &implementation{
		db: db,
	}
}

func (i *implementation) GetUserInterests(ctx context.Context, id domain.UserID) (domain.Interests, error) {
	return i.db.ListUserInterests(ctx, id)
}

func (i *implementation) GetGroupInterests(ctx context.Context, id domain.GroupID) (domain.Interests, error) {
	return i.db.ListGroupInterests(ctx, id)
}
