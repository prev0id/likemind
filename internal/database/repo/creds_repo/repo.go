package creds_repo

import (
	"context"
	"time"

	"likemind/internal/database"
	"likemind/internal/database/model"
	"likemind/internal/domain"

	sql "github.com/huandu/go-sqlbuilder"
)

type DB interface {
	Create(ctx context.Context, creds model.Credentials) error
	UpdatePassword(ctx context.Context, id string, newPassword []byte) error
	GetByLogin(ctx context.Context, login domain.Email) (model.Credentials, error)
	GetByID(ctx context.Context, id string) (model.Credentials, error)
}

type Repo struct{}

func (r *Repo) Create(ctx context.Context, creds model.Credentials) error {
	now := time.Now()

	q := sql.InsertInto(model.TableCredentials)
	q.Cols(
		model.CredentialsID,
		model.CredentialsPassword,
		model.CredentialsLogin,
		model.CredentialsUserID,
		model.CredentialsCreatedAt,
		model.CredentialsUpdatedAt,
	)
	q.Values(
		creds.ID,
		creds.Password,
		creds.Login,
		creds.UserID,
		now,
		now,
	)

	if _, err := database.Exec(ctx, q); err != nil {
		return err
	}

	return nil
}

func (r *Repo) UpdatePassword(ctx context.Context, id string, newPassword []byte) error {
	q := sql.Update(model.TableCredentials)
	q.Set(
		q.Assign(model.CredentialsPassword, newPassword),
		q.Assign(model.CredentialsUpdatedAt, time.Now()),
	)
	q.Where(q.Equal(model.CredentialsID, id))

	if _, err := database.Exec(ctx, q); err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetByLogin(ctx context.Context, login domain.Email) (model.Credentials, error) {
	q := sql.Select(
		model.CredentialsID,
		model.CredentialsPassword,
		model.CredentialsLogin,
		model.CredentialsUserID,
		model.CredentialsCreatedAt,
		model.CredentialsUpdatedAt,
	)
	q.From(model.TableCredentials)
	q.Where(q.Equal(model.CredentialsLogin, login))

	result, err := database.Get[model.Credentials](ctx, q)
	if err != nil {
		return model.Credentials{}, nil
	}

	return result, nil
}

func (r *Repo) GetByID(ctx context.Context, id string) (model.Credentials, error) {
	q := sql.Select(
		model.CredentialsID,
		model.CredentialsPassword,
		model.CredentialsLogin,
		model.CredentialsUserID,
		model.CredentialsCreatedAt,
		model.CredentialsUpdatedAt,
	)
	q.From(model.TableCredentials)
	q.Where(q.Equal(model.CredentialsID, id))

	result, err := database.Get[model.Credentials](ctx, q)
	if err != nil {
		return model.Credentials{}, nil
	}

	return result, nil
}
