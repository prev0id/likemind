package profile

import (
	"context"
	"fmt"

	"likemind/internal/domain"
)

type Service interface {
	CreateUser(ctx context.Context, user domain.User) (int64, error)
	UpdateUser(ctx context.Context, user domain.User) error
	DeleteUser(ctx context.Context, id int64) error
	GetUser(ctx context.Context, id int64) (domain.User, error)
}

type implementation struct {
	db domain.DataProvider[domain.User]
}

func New(provider domain.DataProvider[domain.User]) Service {
	return &implementation{
		db: provider,
	}
}

func (i *implementation) CreateUser(ctx context.Context, user domain.User) (int64, error) {
	_, err := i.db.Get(ctx, domain.FieldNickname, user.Nickname)
	if err == nil {
		return 0, fmt.Errorf("username '%s' already exists", user.Nickname)
	}

	id, err := i.db.Insert(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("i.db.CreateUser: %w", err)
	}

	return id, nil
}

func (i *implementation) UpdateUser(ctx context.Context, user domain.User) error {
	if err := i.db.Update(ctx, user); err != nil {
		return fmt.Errorf("i.db.UpdateUser: %w", err)
	}

	return nil
}

func (i *implementation) GetUser(ctx context.Context, id int64) (domain.User, error) {
	user, err := i.db.Get(ctx, domain.FieldID, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("i.db.GetUser: %w", err)
	}

	return user, nil
}

func (i *implementation) DeleteUser(ctx context.Context, id int64) error {
	if err := i.db.Delete(ctx, id); err != nil {
		return fmt.Errorf("i.db.DeleteUser: %w", err)
	}

	return nil
}
