package interest_adapter

import (
	"context"
	"fmt"

	"likemind/internal/database/model"
	"likemind/internal/database/repo/interest_repo"
	"likemind/internal/domain"
)

type Adapter interface {
	ListInterests(ctx context.Context) ([]domain.InterestGroup, error)

	ListUserInterests(ctx context.Context, id domain.UserID) ([]domain.InterestGroup, error)
	AddInterestToUser(ctx context.Context, userID domain.UserID, interestID domain.InterestID) error
	DeleteInterestFromUser(ctx context.Context, userID domain.UserID, interestID domain.InterestID) error

	ListGroupInterests(ctx context.Context, id domain.GroupID) ([]domain.InterestGroup, error)
	AddInterestToGroup(ctx context.Context, groupID domain.GroupID, interestID domain.InterestID) error
	DeleteInterestFromGroup(ctx context.Context, groupID domain.GroupID, interestID domain.InterestID) error
}

var _ (Adapter) = (*Implementation)(nil)

func New(repo interest_repo.DB) *Implementation {
	return &Implementation{
		repo: repo,
	}
}

type Implementation struct {
	repo interest_repo.DB
}

func (i *Implementation) ListUserInterests(ctx context.Context, id domain.UserID) ([]domain.InterestGroup, error) {
	interests, err := i.repo.ListInterests(ctx)
	if err != nil {
		return nil, fmt.Errorf("i.repo.ListInterests: %w", err)
	}

	groups, err := i.repo.ListInterestGroups(ctx)
	if err != nil {
		return nil, fmt.Errorf("i.repo.ListInterestGroups: %w", err)
	}

	userInterests, err := i.repo.GetUserInterestsByID(ctx, int64(id))
	if err != nil {
		return nil, fmt.Errorf("i.repo.GetUserInterestsByID: %w", err)
	}

	return repoUserInterestsToDomain(userInterests, groups, interests), nil
}

func (i *Implementation) AddInterestToUser(ctx context.Context, userID domain.UserID, interestID domain.InterestID) error {
	err := i.repo.AddInterestToUser(ctx, model.UserInterest{
		UserID:     int64(userID),
		InterestID: int64(interestID),
	})
	if err != nil {
		return fmt.Errorf("i.repo.AddInterestToUser: %w", err)
	}

	return nil
}

func (i *Implementation) DeleteInterestFromUser(ctx context.Context, userID domain.UserID, interestID domain.InterestID) error {
	err := i.repo.RemoveInterestFromUser(ctx, int64(userID), int64(interestID))
	if err != nil {
		return fmt.Errorf("i.repo.RemoveInterestFromUser: %w", err)
	}
	return nil
}

func (i *Implementation) ListGroupInterests(ctx context.Context, id domain.GroupID) ([]domain.InterestGroup, error) {
	interests, err := i.repo.ListInterests(ctx)
	if err != nil {
		return nil, fmt.Errorf("i.repo.ListInterests: %w", err)
	}

	groups, err := i.repo.ListInterestGroups(ctx)
	if err != nil {
		return nil, fmt.Errorf("i.repo.ListInterestGroups: %w", err)
	}

	groupInterests, err := i.repo.GetGroupInterestsByID(ctx, int64(id))
	if err != nil {
		return nil, fmt.Errorf("i.repo.GetUserInterestsByID: %w", err)
	}

	return repoGroupInterestsToDomain(groupInterests, groups, interests), nil
}

func (i *Implementation) AddInterestToGroup(ctx context.Context, groupID domain.GroupID, interestID domain.InterestID) error {
	err := i.repo.AddInterestToGroup(ctx, model.GroupInterest{
		GroupID:    int64(groupID),
		InterestID: int64(interestID),
	})
	if err != nil {
		return fmt.Errorf("i.repo.AddInterestToGroup: %w", err)
	}

	return nil
}

func (i *Implementation) DeleteInterestFromGroup(ctx context.Context, groupID domain.GroupID, interestID domain.InterestID) error {
	err := i.repo.RemoveInterestFromGroup(ctx, int64(groupID), int64(interestID))
	if err != nil {
		return fmt.Errorf("i.repo.RemoveInterestFromGroup: %w", err)
	}
	return nil
}

func (i *Implementation) ListInterests(ctx context.Context) ([]domain.InterestGroup, error) {
	interests, err := i.repo.ListInterests(ctx)
	if err != nil {
		return nil, fmt.Errorf("i.repo.ListInterests: %w", err)
	}

	groups, err := i.repo.ListInterestGroups(ctx)
	if err != nil {
		return nil, fmt.Errorf("i.repo.ListInterestGroups: %w", err)
	}

	return repoUserInterestsToDomain(nil, groups, interests), nil
}
