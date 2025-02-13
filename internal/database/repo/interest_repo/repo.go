package interest_repo

import (
	"context"
	"time"

	"likemind/internal/database"
	"likemind/internal/database/model"

	sql "github.com/huandu/go-sqlbuilder"
)

type DB interface {
	ListInterests(ctx context.Context) ([]model.Interest, error)
	ListInterestsByIDs(ctx context.Context, ids []int64) ([]model.Interest, error)

	GetUserInterestsByID(ctx context.Context, userID int64) ([]model.UserInterest, error)
	AddInterestToUser(ctx context.Context, interest model.UserInterest) error
	RemoveInterestFromUser(ctx context.Context, userID, interestID int64) error

	GetGroupInterestsByID(ctx context.Context, groupID int64) ([]model.GroupInterest, error)
	AddInterestToGroup(ctx context.Context, interest model.GroupInterest) error
	RemoveInterestFromGroup(ctx context.Context, groupID, interestID int64) error
}

var _ DB = (*Repo)(nil)

type Repo struct{}

func (r *Repo) ListInterests(ctx context.Context) ([]model.Interest, error) {
	q := sql.Select(
		model.InterestID,
		model.InterestName,
		model.InterestDescription,
		model.InterestCreatedAt,
		model.InterestUpdatedAt,
	)
	q.From(model.TableInterest)

	return database.Select[model.Interest](ctx, q)
}

func (r *Repo) ListInterestsByIDs(ctx context.Context, ids []int64) ([]model.Interest, error) {
	q := sql.Select(
		model.InterestID,
		model.InterestName,
		model.InterestDescription,
		model.InterestCreatedAt,
		model.InterestUpdatedAt,
	)
	q.From(model.TableInterest)
	q.Where(q.In(model.InterestID, toInterfaceSlice(ids)...))

	return database.Select[model.Interest](ctx, q)
}

func (r *Repo) GetUserInterestsByID(ctx context.Context, userID int64) ([]model.UserInterest, error) {
	q := sql.Select(
		model.UserInterestUserID,
		model.UserInterestInterestID,
		model.UserInterestCreatedAt,
	)
	q.From(model.TableUserInterest)
	q.Where(q.Equal(model.UserInterestUserID, userID))

	return database.Select[model.UserInterest](ctx, q)
}

func (r *Repo) AddInterestToUser(ctx context.Context, interest model.UserInterest) error {
	now := time.Now()
	interest.CreatedAt = now

	q := sql.InsertInto(model.TableUserInterest)
	q.Cols(
		model.UserInterestUserID,
		model.UserInterestInterestID,
		model.UserInterestCreatedAt,
	)
	q.Values(
		interest.UserID,
		interest.InterestID,
		interest.CreatedAt,
	)

	if _, err := database.Exec(ctx, q); err != nil {
		return err
	}

	return nil
}

func (r *Repo) RemoveInterestFromUser(ctx context.Context, userID, interestID int64) error {
	q := sql.DeleteFrom(model.TableUserInterest)
	q.Where(
		q.Equal(model.UserInterestUserID, userID),
		q.Equal(model.UserInterestInterestID, interestID),
	)

	if _, err := database.Exec(ctx, q); err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetGroupInterestsByID(ctx context.Context, groupID int64) ([]model.GroupInterest, error) {
	q := sql.Select(
		model.GroupInterestGroupID,
		model.GroupInterestInterestID,
		model.GroupInterestCreatedAt,
	)
	q.From(model.TableGroupInterest)
	q.Where(q.Equal(model.GroupInterestGroupID, groupID))

	return database.Select[model.GroupInterest](ctx, q)
}

func (r *Repo) AddInterestToGroup(ctx context.Context, interest model.GroupInterest) error {
	now := time.Now()
	interest.CreatedAt = now

	q := sql.InsertInto(model.TableGroupInterest)
	q.Cols(
		model.GroupInterestGroupID,
		model.GroupInterestInterestID,
		model.GroupInterestCreatedAt,
	)
	q.Values(
		interest.GroupID,
		interest.InterestID,
		interest.CreatedAt,
	)

	if _, err := database.Exec(ctx, q); err != nil {
		return err
	}

	return nil
}

func (r *Repo) RemoveInterestFromGroup(ctx context.Context, groupID, interestID int64) error {
	q := sql.DeleteFrom(model.TableGroupInterest)
	q.Where(
		q.Equal(model.GroupInterestGroupID, groupID),
		q.Equal(model.GroupInterestInterestID, interestID),
	)

	if _, err := database.Exec(ctx, q); err != nil {
		return err
	}

	return nil
}

// --- Helper Function ---

// toInterfaceSlice converts a slice of int64 to a slice of interface{}.
func toInterfaceSlice(ints []int64) []interface{} {
	s := make([]interface{}, len(ints))
	for i, v := range ints {
		s[i] = v
	}
	return s
}
