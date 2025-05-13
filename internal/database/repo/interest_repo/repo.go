package interest_repo

import (
	"context"
	"fmt"
	"likemind/internal/database"
	"likemind/internal/database/model"
	"time"

	sql "github.com/huandu/go-sqlbuilder"
)

type DB interface {
	ListInterests(ctx context.Context) ([]model.Interest, error)
	ListInterestGroups(ctx context.Context) ([]model.InterestGroup, error)
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
		model.InterestGroupID,
		model.InterestName,
		model.InterestDescription,
	)
	q.From(model.TableInterests)

	return database.Select[model.Interest](ctx, q)
}

func (r *Repo) ListInterestGroups(ctx context.Context) ([]model.InterestGroup, error) {
	q := sql.Select(
		model.InterestGroupsID,
		model.InterestGroupsName,
	)
	q.From(model.TableInterestGroups)

	return database.Select[model.InterestGroup](ctx, q)
}

func (r *Repo) ListInterestsByIDs(ctx context.Context, ids []int64) ([]model.Interest, error) {
	q := sql.Select(
		model.InterestID,
		model.InterestName,
		model.InterestGroupID,
		model.InterestDescription,
	)
	q.From(model.TableInterests)
	q.Where(q.In(model.InterestID, toInterfaceSlice(ids)...))

	return database.Select[model.Interest](ctx, q)
}

func (r *Repo) GetUserInterestsByID(ctx context.Context, userID int64) ([]model.UserInterest, error) {
	q := sql.Select(
		model.UserInterestUserID,
		model.UserInterestInterestID,
		model.UserInterestCreatedAt,
	)
	q.From(model.TableUserInterests)
	q.Where(q.Equal(model.UserInterestUserID, userID))

	return database.Select[model.UserInterest](ctx, q)
}

func (r *Repo) AddInterestToUser(ctx context.Context, interest model.UserInterest) error {
	now := time.Now()
	interest.CreatedAt = now

	q := sql.InsertInto(model.TableUserInterests)
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
	q := sql.DeleteFrom(model.TableUserInterests)
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
	q.From(model.TableGroupInterests)
	q.Where(q.Equal(model.GroupInterestGroupID, groupID))

	return database.Select[model.GroupInterest](ctx, q)
}

func (r *Repo) AddInterestToGroup(ctx context.Context, interest model.GroupInterest) error {
	now := time.Now()
	interest.CreatedAt = now

	q := sql.InsertInto(model.TableGroupInterests)
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
	q := sql.DeleteFrom(model.TableGroupInterests)
	q.Where(
		q.Equal(model.GroupInterestGroupID, groupID),
		q.Equal(model.GroupInterestInterestID, interestID),
	)

	if _, err := database.Exec(ctx, q); err != nil {
		return err
	}

	return nil
}

func toInterfaceSlice(ints []int64) []any {
	s := make([]any, len(ints))
	for i, v := range ints {
		s[i] = v
	}
	return s
}

func (r *Repo) SearchUsers(ctx context.Context, userInterests, include, exlcude []int64) ([]model.SearchResult, error) {
	sql := `
		SELECT
			user_id AS id,
			count(*) FILTER (WHERE interest_id = ANY($1)) AS commont
		FROM user_interests
		GROUP BY user_id
		HEAVING
			COUNT(DISTINCT interest_id) FILTER (WHERE interest_id = ANY($2)) = cardinality($1)
			AND COUNT(*) (WHERE interest_id = ANY($3)) = 0
		ORDER BY commont DESC;
	`

	result, err := database.Select[model.SearchResult](ctx, database.RawSQL(sql, userInterests, include, exlcude))
	if err != nil {
		return nil, fmt.Errorf("database.Select: %w", err)
	}

	return result, nil
}
