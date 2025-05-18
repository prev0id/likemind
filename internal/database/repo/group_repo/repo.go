package group_repo

import (
	"context"
	"fmt"
	"time"

	"likemind/internal/database"
	"likemind/internal/database/model"

	sql "github.com/huandu/go-sqlbuilder"
)

type DB interface {
	Create(ctx context.Context, group model.Group) (int64, error)
	Update(ctx context.Context, group model.Group) error
	GetByID(ctx context.Context, id int64) (model.Group, error)
	List(ctx context.Context) ([]model.Group, error)
	Delete(ctx context.Context, id int64) error
	ListUserSubscriptions(ctx context.Context, id int64) ([]model.UserSubscription, error)
	AddUserSubscription(ctx context.Context, sub model.UserSubscription) error
	DeleteUserSubscription(ctx context.Context, sub model.UserSubscription) error
}

var _ DB = (*Repo)(nil)

type Repo struct{}

func (r *Repo) Create(ctx context.Context, group model.Group) (int64, error) {
	now := time.Now()
	group.CreatedAt = now
	group.UpdatedAt = now

	q := sql.InsertInto(model.TableGroups)
	q.Cols(
		model.GroupID,
		model.GroupName,
		model.GroupDescription,
		model.GroupAuthorID,
		model.GroupCreatedAt,
		model.GroupUpdatedAt,
	)
	q.Values(
		group.ID,
		group.Name,
		group.Description,
		group.AuthorID,
		group.CreatedAt,
		group.UpdatedAt,
	)
	q.SQL("RETURNING " + model.GroupID)

	id, err := database.Get[int64](ctx, q)
	if err != nil {
		return 0, fmt.Errorf("database.Get: %w", err)
	}

	return id, nil
}

func (r *Repo) Update(ctx context.Context, group model.Group) error {
	group.UpdatedAt = time.Now()

	q := sql.Update(model.TableGroups)
	q.Set(
		q.Assign(model.GroupName, group.Name),
		q.Assign(model.GroupDescription, group.Description),
		q.Assign(model.GroupAuthorID, group.AuthorID),
		q.Assign(model.GroupUpdatedAt, group.UpdatedAt),
	)
	q.Where(q.Equal(model.GroupID, group.ID))

	if _, err := database.Exec(ctx, q); err != nil {
		return fmt.Errorf("database.Exec: %w", err)
	}
	return nil
}

func (r *Repo) GetByID(ctx context.Context, id int64) (model.Group, error) {
	q := sql.Select(
		model.GroupID,
		model.GroupName,
		model.GroupDescription,
		model.GroupAuthorID,
		model.GroupCreatedAt,
		model.GroupUpdatedAt,
	)
	q.From(model.TableGroups)
	q.Where(q.Equal(model.GroupID, id))

	result, err := database.Get[model.Group](ctx, q)
	if err != nil {
		return model.Group{}, fmt.Errorf("database.Get: %w", err)
	}

	return result, nil
}

func (r *Repo) List(ctx context.Context) ([]model.Group, error) {
	q := sql.Select(
		model.GroupID,
		model.GroupName,
		model.GroupDescription,
		model.GroupAuthorID,
		model.GroupCreatedAt,
		model.GroupUpdatedAt,
	)
	q.From(model.TableGroups)

	results, err := database.Select[model.Group](ctx, q)
	if err != nil {
		return nil, fmt.Errorf("database.Select: %w", err)
	}

	return results, nil
}

func (r *Repo) Delete(ctx context.Context, id int64) error {
	q := sql.DeleteFrom(model.TableGroups)
	q.Where(q.Equal(model.GroupID, id))

	if _, err := database.Exec(ctx, q); err != nil {
		return fmt.Errorf("database.Exec: %w", err)
	}

	return nil
}

func (r *Repo) ListUserSubscriptions(ctx context.Context, id int64) ([]model.UserSubscription, error) {
	q := sql.Select(
		model.UserSubscriptionUserID,
		model.UserSubscriptionGroupID,
		model.UserSubscriptionCreateAt,
	)
	q.From(model.TableUserSubscriptions)

	results, err := database.Select[model.UserSubscription](ctx, q)
	if err != nil {
		return nil, fmt.Errorf("database.Select: %w", err)
	}

	return results, nil
}

func (r *Repo) AddUserSubscription(ctx context.Context, sub model.UserSubscription) error {
	q := sql.InsertInto(model.TableUserSubscriptions)
	q.Cols(
		model.UserSubscriptionUserID,
		model.UserSubscriptionGroupID,
		model.UserSubscriptionCreateAt,
	)
	q.Values(
		sub.UserID,
		sub.GroupID,
		time.Now(),
	)

	if _, err := database.Exec(ctx, q); err != nil {
		return fmt.Errorf("database.Exec: %w", err)
	}

	return nil
}

func (r *Repo) DeleteUserSubscription(ctx context.Context, sub model.UserSubscription) error {
	q := sql.DeleteFrom(model.TableUserSubscriptions)
	q.Where(
		q.Equal(model.UserSubscriptionUserID, sub.UserID),
		q.Equal(model.UserSubscriptionGroupID, sub.GroupID),
	)

	if _, err := database.Exec(ctx, q); err != nil {
		return fmt.Errorf("database.Exec: %w", err)
	}
	return nil
}
