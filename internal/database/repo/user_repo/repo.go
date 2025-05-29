package user_repo

import (
	"context"
	"time"

	"likemind/internal/database"
	"likemind/internal/database/model"

	sql "github.com/huandu/go-sqlbuilder"
)

type DB interface {
	Create(ctx context.Context, user model.User) (int64, error)
	Update(ctx context.Context, user model.User) error
	GetByID(ctx context.Context, id int64) (model.User, error)
	GetByEmail(ctx context.Context, email string) (model.User, error)
	GetByUsername(ctx context.Context, username string) (model.User, error)
	List(ctx context.Context) ([]model.User, error)
	Delete(ctx context.Context, userID int64) error
}

var _ DB = (*Repo)(nil)

type Repo struct{}

func (r *Repo) Create(ctx context.Context, user model.User) (int64, error) {
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	q := sql.InsertInto(model.TableUsers)
	q.Cols(
		model.UserNickname,
		model.UserName,
		model.UserSurname,
		model.UserAbout,
		model.UserEmail,
		model.UserPassword,
		model.UserLocation,
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
		user.Location,
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

func (r *Repo) Update(ctx context.Context, user model.User) error {
	user.UpdatedAt = time.Now()

	q := sql.Update(model.TableUsers)
	q.Set(
		q.Assign(model.UserNickname, user.Nickname),
		q.Assign(model.UserName, user.Name),
		q.Assign(model.UserSurname, user.Surname),
		q.Assign(model.UserAbout, user.About),
		q.Assign(model.UserEmail, user.Email),
		q.Assign(model.UserPassword, user.Password),
		q.Assign(model.UserLocation, user.Location),
		q.Assign(model.UserUpdatedAt, user.UpdatedAt),
	)
	q.Where(q.Equal(model.UserID, user.ID))

	if _, err := database.Exec(ctx, q); err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetByID(ctx context.Context, id int64) (model.User, error) {
	q := sql.Select(
		model.UserID,
		model.UserNickname,
		model.UserName,
		model.UserSurname,
		model.UserAbout,
		model.UserEmail,
		model.UserPassword,
		model.UserLocation,
		model.UserCreatedAt,
		model.UserUpdatedAt,
	)
	q.From(model.TableUsers)
	q.Where(q.Equal(model.UserID, id))

	result, err := database.Get[model.User](ctx, q)
	if err != nil {
		return model.User{}, err
	}

	return result, nil
}

func (r *Repo) GetByUsername(ctx context.Context, username string) (model.User, error) {
	q := sql.Select(
		model.UserID,
		model.UserNickname,
		model.UserName,
		model.UserSurname,
		model.UserAbout,
		model.UserEmail,
		model.UserPassword,
		model.UserLocation,
		model.UserCreatedAt,
		model.UserUpdatedAt,
	)
	q.From(model.TableUsers)
	q.Where(q.Equal(model.UserNickname, username))

	result, err := database.Get[model.User](ctx, q)
	if err != nil {
		return model.User{}, err
	}

	return result, nil
}

func (r *Repo) List(ctx context.Context) ([]model.User, error) {
	q := sql.Select(
		model.UserID,
		model.UserNickname,
		model.UserName,
		model.UserSurname,
		model.UserAbout,
		model.UserEmail,
		model.UserPassword,
		model.UserLocation,
		model.UserCreatedAt,
		model.UserUpdatedAt,
	)
	q.From(model.TableUsers)

	result, err := database.Select[model.User](ctx, q)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repo) Delete(ctx context.Context, userID int64) error {
	q := sql.DeleteFrom(model.TableUsers)
	q.Where(q.Equal(model.UserID, userID))

	if _, err := database.Exec(ctx, q); err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetByEmail(ctx context.Context, email string) (model.User, error) {
	q := sql.Select(
		model.UserID,
		model.UserNickname,
		model.UserName,
		model.UserSurname,
		model.UserAbout,
		model.UserEmail,
		model.UserPassword,
		model.UserLocation,
		model.UserCreatedAt,
		model.UserUpdatedAt,
	)
	q.From(model.TableUsers)
	q.Where(q.Equal(model.UserEmail, email))

	result, err := database.Get[model.User](ctx, q)
	if err != nil {
		return model.User{}, err
	}

	return result, nil
}
