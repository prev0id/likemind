package comment_repo

import (
	"context"
	"fmt"
	"time"

	"likemind/internal/database"
	"likemind/internal/database/model"

	sql "github.com/huandu/go-sqlbuilder"
)

type DB interface {
	Create(ctx context.Context, comment model.Comment) (int64, error)
	Update(ctx context.Context, comment model.Comment) error
	GetByID(ctx context.Context, id int64) (model.Comment, error)
	ListByPost(ctx context.Context, postID int64) ([]model.Comment, error)
	Delete(ctx context.Context, id int64) error
}

var _ DB = (*Repo)(nil)

type Repo struct{}

func (r *Repo) Create(ctx context.Context, comment model.Comment) (int64, error) {
	now := time.Now()
	comment.CreatedAt = now
	comment.UpdatedAt = now

	q := sql.InsertInto(model.TableComments)
	q.Cols(
		model.CommentID,
		model.CommentPostID,
		model.CommentContent,
		model.CommentAuthorID,
		model.CommentCreatedAt,
		model.CommentUpdatedAt,
	)
	q.Values(
		comment.ID,
		comment.PostID,
		comment.Content,
		comment.AuthorID,
		comment.CreatedAt,
		comment.UpdatedAt,
	)
	q.SQL("RETURNING " + model.CommentID)

	id, err := database.Get[int64](ctx, q)
	if err != nil {
		return 0, fmt.Errorf("database.Get: %w", err)
	}

	return id, nil
}

func (r *Repo) Update(ctx context.Context, comment model.Comment) error {
	comment.UpdatedAt = time.Now()

	q := sql.Update(model.TableComments)
	q.Set(
		q.Assign(model.CommentContent, comment.Content),
		q.Assign(model.CommentPostID, comment.PostID),
		q.Assign(model.CommentAuthorID, comment.AuthorID),
		q.Assign(model.CommentUpdatedAt, comment.UpdatedAt),
	)
	q.Where(
		q.Equal(model.CommentID, comment.ID),
	)

	if _, err := database.Exec(ctx, q); err != nil {
		return fmt.Errorf("database.Exec: %w", err)
	}
	return nil
}

func (r *Repo) GetByID(ctx context.Context, id int64) (model.Comment, error) {
	q := sql.Select(
		model.CommentID,
		model.CommentPostID,
		model.CommentContent,
		model.CommentAuthorID,
		model.CommentCreatedAt,
		model.CommentUpdatedAt,
	)
	q.From(model.TableComments)
	q.Where(
		q.Equal(model.CommentID, id),
	)

	comment, err := database.Get[model.Comment](ctx, q)
	if err != nil {
		return model.Comment{}, fmt.Errorf("database.Get: %w", err)
	}

	return comment, nil
}

func (r *Repo) ListByPost(ctx context.Context, postID int64) ([]model.Comment, error) {
	q := sql.Select(
		model.CommentID,
		model.CommentPostID,
		model.CommentContent,
		model.CommentAuthorID,
		model.CommentCreatedAt,
		model.CommentUpdatedAt,
	)
	q.From(model.TableComments)
	q.Where(
		q.Equal(model.CommentPostID, postID),
	)

	comments, err := database.Select[model.Comment](ctx, q)
	if err != nil {
		return nil, fmt.Errorf("database.Select: %w", err)
	}

	return comments, nil
}

func (r *Repo) Delete(ctx context.Context, id int64) error {
	q := sql.DeleteFrom(model.TableComments)
	q.Where(q.Equal(model.CommentID, id))

	if _, err := database.Exec(ctx, q); err != nil {
		return fmt.Errorf("database.Exec: %w", err)
	}

	return nil
}
