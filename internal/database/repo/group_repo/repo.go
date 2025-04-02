package group_repo

import (
	"context"
	"time"

	"likemind/internal/database"
	"likemind/internal/database/model"

	sql "github.com/huandu/go-sqlbuilder"
)

type DB interface {
	CreateGroup(ctx context.Context, group model.Group) error
	UpdateGroup(ctx context.Context, group model.Group) error
	GetGroupByAlias(ctx context.Context, alias string) (model.Group, error)
	GetGroupByID(ctx context.Context, id int64) (model.Group, error)
	ListGroups(ctx context.Context) ([]model.Group, error)
}

var _ DB = (*Repo)(nil)

type Repo struct{}

func (r *Repo) CreateGroup(ctx context.Context, group model.Group) error {
	now := time.Now()
	group.CreatedAt = now
	group.UpdatedAt = now

	q := sql.InsertInto(model.TableGroups)
	q.Cols(
		model.GroupID,
		model.GroupPictureID,
		model.GroupName,
		model.GroupAlias,
		model.GroupDescription,
		model.GroupAuthorID,
		model.GroupCreatedAt,
		model.GroupUpdatedAt,
	)
	q.Values(
		group.ID,
		group.PictureID,
		group.Name,
		group.Alias,
		group.Description,
		group.AuthorID,
		group.CreatedAt,
		group.UpdatedAt,
	)

	if _, err := database.Exec(ctx, q); err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateGroup(ctx context.Context, group model.Group) error {
	group.UpdatedAt = time.Now()

	q := sql.Update(model.TableGroups)
	q.Set(
		q.Assign(model.GroupPictureID, group.PictureID),
		q.Assign(model.GroupName, group.Name),
		q.Assign(model.GroupAlias, group.Alias),
		q.Assign(model.GroupDescription, group.Description),
		q.Assign(model.GroupAuthorID, group.AuthorID),
		q.Assign(model.GroupUpdatedAt, group.UpdatedAt),
	)
	q.Where(q.Equal(model.GroupID, group.ID))

	if _, err := database.Exec(ctx, q); err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetGroupByAlias(ctx context.Context, alias string) (model.Group, error) {
	q := sql.Select(
		model.GroupID,
		model.GroupPictureID,
		model.GroupName,
		model.GroupAlias,
		model.GroupDescription,
		model.GroupAuthorID,
		model.GroupCreatedAt,
		model.GroupUpdatedAt,
	)
	q.From(model.TableGroups)
	q.Where(q.Equal(model.GroupAlias, alias))

	result, err := database.Get[model.Group](ctx, q)
	if err != nil {
		return model.Group{}, err
	}

	return result, nil
}

func (r *Repo) GetGroupByID(ctx context.Context, id int64) (model.Group, error) {
	q := sql.Select(
		model.GroupID,
		model.GroupPictureID,
		model.GroupName,
		model.GroupAlias,
		model.GroupDescription,
		model.GroupAuthorID,
		model.GroupCreatedAt,
		model.GroupUpdatedAt,
	)
	q.From(model.TableGroups)
	q.Where(q.Equal(model.GroupID, id))

	result, err := database.Get[model.Group](ctx, q)
	if err != nil {
		return model.Group{}, err
	}

	return result, nil
}

func (r *Repo) ListGroups(ctx context.Context) ([]model.Group, error) {
	q := sql.Select(
		model.GroupID,
		model.GroupPictureID,
		model.GroupName,
		model.GroupAlias,
		model.GroupDescription,
		model.GroupAuthorID,
		model.GroupCreatedAt,
		model.GroupUpdatedAt,
	)
	q.From(model.TableGroups)

	results, err := database.Select[model.Group](ctx, q)
	if err != nil {
		return nil, err
	}

	return results, nil
}
