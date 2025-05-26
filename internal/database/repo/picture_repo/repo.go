package profile_picture_repo

import (
	"context"
	"time"

	"likemind/internal/database"
	"likemind/internal/database/model"

	sql "github.com/huandu/go-sqlbuilder"
)

type DB interface {
	AddProfilePicture(ctx context.Context, picture model.ProfilePicture) error
	GetProfilePicturesByUserID(ctx context.Context, userID int64) ([]model.ProfilePicture, error)
	RemovePictureByID(ctx context.Context, id string) error
}

var _ DB = (*Repo)(nil)

type Repo struct{}

func (r *Repo) AddProfilePicture(ctx context.Context, picture model.ProfilePicture) error {
	now := time.Now()
	picture.CreatedAt = now
	picture.UpdatedAt = now

	q := sql.InsertInto(model.TableProfilePictures)
	q.Cols(
		model.ProfilePictureID,
		model.ProfilePictureUserID,
		model.ProfilePictureCreatedAt,
		model.ProfilePictureUpdatedAt,
	)
	q.Values(
		picture.ID,
		picture.UserID,
		picture.CreatedAt,
		picture.UpdatedAt,
	)

	if _, err := database.Exec(ctx, q); err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetProfilePicturesByUserID(ctx context.Context, userID int64) ([]model.ProfilePicture, error) {
	q := sql.Select(
		model.ProfilePictureID,
		model.ProfilePictureUserID,
		model.ProfilePictureCreatedAt,
		model.ProfilePictureUpdatedAt,
	)
	q.From(model.TableProfilePictures)
	q.Where(q.Equal(model.ProfilePictureUserID, userID))
	q.Desc().OrderBy(model.ProfilePictureUpdatedAt)

	results, err := database.Select[model.ProfilePicture](ctx, q)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r *Repo) RemovePictureByID(ctx context.Context, id string) error {
	q := sql.DeleteFrom(model.TableProfilePictures)
	q.Where(q.Equal(model.ProfilePictureID, id))

	if _, err := database.Exec(ctx, q); err != nil {
		return err
	}

	return nil
}
