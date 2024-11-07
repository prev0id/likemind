package profile

import (
	"context"
	"fmt"

	"likemind/internal/domain"
)

type Service interface {
	UpdateName(ctx context.Context, id uint64, newName string) error
}

type db interface {
	ListUsers(ctx context.Context) ([]domain.User, error)
	GetUser(ctx context.Context, id uint64) (domain.User, error)
	UpdateUser(ctx context.Context, user domain.User) error
	CreateUser(ctx context.Context, name, username string) (uint64, error)
}

type implementation struct {
	db db
}

func New(db db) Service {
	return &implementation{}
}

func (i *implementation) RegestryNewUser(ctx context.Context, name, username string) (uint64, error) {
	id, err := i.db.CreateUser(ctx, name, username)
	if err != nil {
		return 0, fmt.Errorf("i.db.CreateUser: %w", err)
	}

	return id, nil
}

func (i *implementation) UpdateName(ctx context.Context, id uint64, name string) error {
	user, err := i.db.GetUser(ctx, id)
	if err != nil {
		return fmt.Errorf("i.db.GetUser: %w", err)
	}

	user.Name = name

	if err := i.db.UpdateUser(ctx, user); err != nil {
		return fmt.Errorf("i.db.UpdateUser: %w", err)
	}

	return nil
}

func (i *implementation) UpdateUsername(ctx context.Context, id uint64, username string) error {
	user, err := i.db.GetUser(ctx, id)
	if err != nil {
		return fmt.Errorf("i.db.GetUser: %w", err)
	}

	user.Username = username

	if err := i.db.UpdateUser(ctx, user); err != nil {
		return fmt.Errorf("i.db.UpdateUser: %w", err)
	}

	return nil
}

func (i *implementation) UpdateAbout(ctx context.Context, id uint64, about string) error {
	user, err := i.db.GetUser(ctx, id)
	if err != nil {
		return fmt.Errorf("i.db.GetUser: %w", err)
	}

	user.About = about

	if err := i.db.UpdateUser(ctx, user); err != nil {
		return fmt.Errorf("i.db.UpdateUser: %w", err)
	}

	return nil
}
