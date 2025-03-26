package session_adapter

import (
	"context"
	"errors"
	"fmt"

	"likemind/internal/database/repo/session_repo"
	"likemind/internal/domain"

	"github.com/jackc/pgx/v5"
)

type Adapter interface {
	GetByToken(ctx context.Context, token domain.SessionToken) (domain.Session, error)
	Create(ctx context.Context, session domain.Session) error
	ClearOld(ctx context.Context) error
	Invalidate(ctx context.Context, id domain.UserID) error
}

type implementation struct {
	repo session_repo.DB
}

func NewAdapter(repo session_repo.DB) Adapter {
	return &implementation{
		repo: repo,
	}
}

func (i *implementation) GetByToken(ctx context.Context, token domain.SessionToken) (domain.Session, error) {
	session, err := i.repo.GetByToken(ctx, string(token))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = domain.ErrNotFound
		}
		return domain.Session{}, fmt.Errorf("i.repo.GetByToken: %w", err)
	}

	return modelSessionToDomain(session), nil
}

func (i *implementation) Create(ctx context.Context, session domain.Session) error {
	err := i.repo.Create(ctx, domainSessionToModel(session))
	if err != nil {
		return fmt.Errorf("i.repo.Create: %w", err)
	}

	return nil
}

func (i *implementation) ClearOld(ctx context.Context) error {
	err := i.repo.ClearOld(ctx)
	if err != nil {
		return fmt.Errorf("i.repo.ClearOld: %w", err)
	}

	return nil
}

func (i *implementation) Invalidate(ctx context.Context, id domain.UserID) error {
	err := i.repo.InvalidateByUserID(ctx, int64(id))
	if err != nil {
		return fmt.Errorf("i.repo.Invalidate: %w", err)
	}

	return nil
}
