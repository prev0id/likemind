package session_repo

import (
	"context"
	"fmt"
	"time"

	"likemind/internal/database"
	"likemind/internal/database/model"

	sql "github.com/huandu/go-sqlbuilder"
)

type DB interface {
	GetByToken(ctx context.Context, token string) (model.Session, error)
	Create(ctx context.Context, session model.Session) error
	ClearOld(ctx context.Context) error
	InvalidateByUserID(ctx context.Context, id int64) error
}

var _ DB = (*Repo)(nil)

type Repo struct{}

func (r *Repo) GetByToken(ctx context.Context, token string) (model.Session, error) {
	q := sql.Select(
		model.CredentialsToken,
		model.CredentialsUserID,
		model.CredentialsCreatedAt,
		model.CredentialsExpiresAt,
	)
	q.From(model.TableSessions)
	q.Where(q.Equal(model.CredentialsToken, token))

	result, err := database.Get[model.Session](ctx, q)
	if err != nil {
		return model.Session{}, fmt.Errorf("database.Get: %w", err)
	}

	return result, nil
}

func (r *Repo) Create(ctx context.Context, session model.Session) error {
	now := time.Now()

	q := sql.InsertInto(model.TableSessions)
	q.Cols(
		model.CredentialsToken,
		model.CredentialsUserID,
		model.CredentialsCreatedAt,
		model.CredentialsExpiresAt,
	)
	q.Values(
		session.Token,
		session.UserID,
		now,
		session.ExpiresAt,
	)

	fmt.Println(q.Build())

	if _, err := database.Exec(ctx, q); err != nil {
		return fmt.Errorf("database.Exec: %w", err)
	}

	return nil
}

func (r *Repo) ClearOld(ctx context.Context) error {
	q := sql.DeleteFrom(model.TableSessions)
	q.Where(q.LessThan(model.CredentialsExpiresAt, time.Now()))

	if _, err := database.Exec(ctx, q); err != nil {
		return fmt.Errorf("database.Exec: %w", err)
	}

	return nil
}

func (r *Repo) InvalidateByUserID(ctx context.Context, id int64) error {
	q := sql.DeleteFrom(model.TableSessions)
	q.Where(q.Equal(model.CredentialsUserID, id))

	if _, err := database.Exec(ctx, q); err != nil {
		return fmt.Errorf("database.Exec: %w", err)
	}

	return nil
}
