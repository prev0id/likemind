package post_repo

import (
	"context"
	"fmt"
	"time"

	"likemind/internal/database"
	"likemind/internal/database/model"

	sql "github.com/huandu/go-sqlbuilder"
)

type DB interface {
	Create(ctx context.Context, post model.Post) (int64, error)
	Update(ctx context.Context, post model.Post) error
	GetByID(ctx context.Context, id int64) (model.Post, error)
	ListByGroupID(ctx context.Context, groupID int64) ([]model.Post, error)
	Delete(ctx context.Context, id int64) error
}

var _ DB = (*Repo)(nil)

type Repo struct{}

func (r *Repo) Create(ctx context.Context, post model.Post) (int64, error) {
	now := time.Now()
	post.CreatedAt = now
	post.UpdatedAt = now

	q := sql.InsertInto(model.TablePosts)
	q.Cols(
		model.PostID,
		model.PostGroupID,
		model.PostContent,
		model.PostAuthorID,
		model.PostCreatedAt,
		model.PostUpdatedAt,
	)
	q.Values(
		post.ID,
		post.GroupID,
		post.Content,
		post.AuthorID,
		post.CreatedAt,
		post.UpdatedAt,
	)
	q.SQL("RETURNING " + model.PostID)

	id, err := database.Get[int64](ctx, q)
	if err != nil {
		return 0, fmt.Errorf("database.Get: %w", err)
	}

	return id, nil
}

func (r *Repo) Update(ctx context.Context, post model.Post) error {
	post.UpdatedAt = time.Now()

	q := sql.Update(model.TablePosts)
	q.Set(
		q.Assign(model.PostContent, post.Content),
		q.Assign(model.PostGroupID, post.GroupID),
		q.Assign(model.PostAuthorID, post.AuthorID),
		q.Assign(model.PostUpdatedAt, post.UpdatedAt),
	)
	q.Where(q.Equal(model.PostID, post.ID))

	if _, err := database.Exec(ctx, q); err != nil {
		return fmt.Errorf("database.Exec: %w", err)
	}
	return nil
}

func (r *Repo) GetByID(ctx context.Context, id int64) (model.Post, error) {
	q := sql.Select(
		model.PostID,
		model.PostGroupID,
		model.PostContent,
		model.PostAuthorID,
		model.PostCreatedAt,
		model.PostUpdatedAt,
	)
	q.From(model.TablePosts)
	q.Where(
		q.Equal(model.PostID, id),
	)

	post, err := database.Get[model.Post](ctx, q)
	if err != nil {
		return model.Post{}, fmt.Errorf("database.Get: %w", err)
	}

	return post, nil
}

func (r *Repo) ListByGroupID(ctx context.Context, groupID int64) ([]model.Post, error) {
	q := sql.Select(
		model.PostID,
		model.PostGroupID,
		model.PostContent,
		model.PostAuthorID,
		model.PostCreatedAt,
		model.PostUpdatedAt,
	)
	q.From(model.TablePosts)
	q.Where(
		q.Equal(model.PostGroupID, groupID),
	)

	posts, err := database.Select[model.Post](ctx, q)
	if err != nil {
		return nil, fmt.Errorf("database.Select: %w", err)
	}

	return posts, nil
}

func (r *Repo) Delete(ctx context.Context, id int64) error {
	q := sql.DeleteFrom(model.TablePosts)
	q.Where(q.Equal(model.PostID, id))

	if _, err := database.Exec(ctx, q); err != nil {
		return fmt.Errorf("database.Exec: %w", err)
	}

	return nil
}
