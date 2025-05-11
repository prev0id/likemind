package interest_adapter

import (
	"context"
	"fmt"

	"likemind/internal/database/repo/interest_repo"
	"likemind/internal/domain"
)

type Adapter interface {
	ListUserInterests(ctx context.Context, id domain.UserID) (domain.Interests, error)
	ListGroupInterests(ctx context.Context, id domain.GroupID) (domain.Interests, error)
	ListInterests(ctx context.Context) (domain.Interests, error)
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

func (i *Implementation) ListUserInterests(ctx context.Context, id domain.UserID) (domain.Interests, error) {
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

func (i *Implementation) ListGroupInterests(ctx context.Context, id domain.GroupID) (domain.Interests, error) {
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

func (i *Implementation) ListInterests(ctx context.Context) (domain.Interests, error) {
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
