package user_repo

import (
	"context"
	"time"

	"likemind/internal/database"
	"likemind/internal/database/model"

	sql "github.com/huandu/go-sqlbuilder"
)

type DB interface {
	CreateUser(ctx context.Context, user model.User) (int64, error)
	UpdateUser(ctx context.Context, user model.User) error
	GetUserByID(ctx context.Context, id int64) (model.User, error)
	ListUsers(ctx context.Context) ([]model.User, error)
	RemoveUser(ctx context.Context, userID int64) error
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
}

var _ DB = (*Repo)(nil)

type Repo struct{}

func (r *Repo) CreateUser(ctx context.Context, user model.User) (int64, error) {
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	q := sql.InsertInto(model.TableUser)
	q.Cols(
		model.UserNickname,
		model.UserName,
		model.UserSurname,
		model.UserAbout,
		model.UserEmail,
		model.UserPassword,
		model.UserCreatedAt,
		model.UserUpdatedAt,
	)
	q.Values(
		user.Nickname,
		user.Name,
		user.Surname,
		user.About,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	)
	q.SQL("RETURNING " + model.UserID)

	id, err := database.Get[int64](ctx, q)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repo) UpdateUser(ctx context.Context, user model.User) error {
	user.UpdatedAt = time.Now()

	q := sql.Update(model.TableUser)
	q.Set(
		q.Assign(model.UserNickname, user.Nickname),
		q.Assign(model.UserName, user.Name),
		q.Assign(model.UserSurname, user.Surname),
		q.Assign(model.UserAbout, user.About),
		q.Assign(model.UserEmail, user.Email),
		q.Assign(model.UserPassword, user.Password),
		q.Assign(model.UserUpdatedAt, user.UpdatedAt),
	)
	q.Where(q.Equal(model.UserID, user.ID))

	if _, err := database.Exec(ctx, q); err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetUserByID(ctx context.Context, id int64) (model.User, error) {
	q := sql.Select(
		model.UserID,
		model.UserNickname,
		model.UserName,
		model.UserSurname,
		model.UserAbout,
		model.UserEmail,
		model.UserPassword,
		model.UserCreatedAt,
		model.UserUpdatedAt,
	)
	q.From(model.TableUser)
	q.Where(q.Equal(model.UserID, id))

	result, err := database.Get[model.User](ctx, q)
	if err != nil {
		return model.User{}, err
	}

	return result, nil
}

func (r *Repo) ListUsers(ctx context.Context) ([]model.User, error) {
	q := sql.Select(
		model.UserID,
		model.UserNickname,
		model.UserName,
		model.UserSurname,
		model.UserAbout,
		model.UserEmail,
		model.UserPassword,
		model.UserCreatedAt,
		model.UserUpdatedAt,
	)
	q.From(model.TableUser)

	result, err := database.Select[model.User](ctx, q)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repo) RemoveUser(ctx context.Context, userID int64) error {
	q := sql.DeleteFrom(model.TableUser)
	q.Where(q.Equal(model.UserID, userID))

	if _, err := database.Exec(ctx, q); err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	q := sql.Select(
		model.UserID,
		model.UserNickname,
		model.UserName,
		model.UserSurname,
		model.UserAbout,
		model.UserEmail,
		model.UserPassword,
		model.UserCreatedAt,
		model.UserUpdatedAt,
	)
	q.From(model.TableUser)
	q.Where(q.Equal(model.UserEmail, email))

	result, err := database.Get[model.User](ctx, q)
	if err != nil {
		return model.User{}, err
	}

	return result, nil
}
