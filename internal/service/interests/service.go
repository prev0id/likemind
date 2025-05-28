package interests

import (
	"context"
	"fmt"

	"likemind/internal/database/adapter/interest_adapter"
	"likemind/internal/domain"
)

type Service interface {
	GetUserInterests(ctx context.Context, id domain.UserID) ([]domain.InterestGroup, error)
	AddInterestToUser(ctx context.Context, userID domain.UserID, interestID domain.InterestID) ([]domain.InterestGroup, error)
	DeleteInterestFromUser(ctx context.Context, userID domain.UserID, interestID domain.InterestID) ([]domain.InterestGroup, error)

	GetGroupInterests(ctx context.Context, id domain.GroupID) ([]domain.InterestGroup, error)
	AddInterestToGroup(ctx context.Context, groupID domain.GroupID, interestID domain.InterestID) ([]domain.InterestGroup, error)
	DeleteInterestFromGroup(ctx context.Context, groupID domain.GroupID, interestID domain.InterestID) ([]domain.InterestGroup, error)

	SearchUsers(ctx context.Context, userID domain.UserID, include, exlcude []int64) ([]domain.UserID, error)
	SearchGroups(ctx context.Context, userID domain.UserID, include, exlcude []int64) ([]domain.GroupID, error)

	ListInterests(ctx context.Context) ([]domain.InterestGroup, error)
}

type implementation struct {
	db interest_adapter.Adapter
}

func New(db interest_adapter.Adapter) Service {
	return &implementation{
		db: db,
	}
}

func (i *implementation) GetUserInterests(ctx context.Context, id domain.UserID) ([]domain.InterestGroup, error) {
	interests, err := i.db.ListUserInterests(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("i.db.ListUserInterests: %w", err)
	}

	return interests, nil
}

func (i *implementation) ListInterests(ctx context.Context) ([]domain.InterestGroup, error) {
	interests, err := i.db.ListInterests(ctx)
	if err != nil {
		return nil, fmt.Errorf("i.db.ListInterests: %w", err)
	}

	return interests, nil
}

func (i *implementation) AddInterestToUser(ctx context.Context, userID domain.UserID, interestID domain.InterestID) ([]domain.InterestGroup, error) {
	if err := i.db.AddInterestToUser(ctx, userID, interestID); err != nil {
		return nil, fmt.Errorf("i.db.AddInterestToUser: %w", err)
	}

	interests, err := i.db.ListUserInterests(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("i.db.ListUserInterests: %w", err)
	}

	return interests, nil
}

func (i *implementation) DeleteInterestFromUser(ctx context.Context, userID domain.UserID, interestID domain.InterestID) ([]domain.InterestGroup, error) {
	if err := i.db.DeleteInterestFromUser(ctx, userID, interestID); err != nil {
		return nil, fmt.Errorf("i.db.DeleteInterestFromUser: %w", err)
	}

	interests, err := i.db.ListUserInterests(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("i.db.ListUserInterests: %w", err)
	}

	return interests, nil
}

func (i *implementation) AddInterestToGroup(ctx context.Context, groupID domain.GroupID, interestID domain.InterestID) ([]domain.InterestGroup, error) {
	if err := i.db.AddInterestToGroup(ctx, groupID, interestID); err != nil {
		return nil, fmt.Errorf("i.db.AddInterestToGroup: %w", err)
	}

	interests, err := i.db.ListGroupInterests(ctx, groupID)
	if err != nil {
		return nil, fmt.Errorf("i.db.ListGroupInterests: %w", err)
	}

	return interests, nil
}

func (i *implementation) DeleteInterestFromGroup(ctx context.Context, groupID domain.GroupID, interestID domain.InterestID) ([]domain.InterestGroup, error) {
	if err := i.db.DeleteInterestFromGroup(ctx, groupID, interestID); err != nil {
		return nil, fmt.Errorf("i.db.DeleteInterestFromGroup: %w", err)
	}

	interests, err := i.db.ListGroupInterests(ctx, groupID)
	if err != nil {
		return nil, fmt.Errorf("i.db.ListGroupInterests: %w", err)
	}

	return interests, nil
}

func (i *implementation) GetGroupInterests(ctx context.Context, id domain.GroupID) ([]domain.InterestGroup, error) {
	return i.db.ListGroupInterests(ctx, id)
}

func (i *implementation) SearchGroups(ctx context.Context, userID domain.UserID, include, exlcude []int64) ([]domain.GroupID, error) {
	return i.db.SearchGroups(ctx, userID, include, exlcude)
}

func (i *implementation) SearchUsers(ctx context.Context, userID domain.UserID, include, exlcude []int64) ([]domain.UserID, error) {
	searched, err := i.db.SearchUsers(ctx, userID, include, exlcude)
	if err != nil {
		return nil, err
	}

	filtered := make([]domain.UserID, 0, len(searched))
	for _, id := range searched {
		if id != userID {
			filtered = append(filtered, id)
		}
	}

	return filtered, nil
}
