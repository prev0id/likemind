package group

import (
	"context"
	"fmt"
	"likemind/internal/common"
	"likemind/internal/domain"
)

func (i *Implementation) CreateGroup(ctx context.Context, group domain.Group) (domain.GroupID, error) {
	group.Author = common.UserIDFromContext(ctx)
	if group.Name != "" {
		return 0, fmt.Errorf("%w: name must not be empty", domain.ErrValidationFailed)
	}

	id, err := i.adapter.CreateGroup(ctx, group)
	if err != nil {
		return 0, fmt.Errorf("i.adapter.CreateGroup: %w", err)
	}

	return id, nil
}

func (s *Implementation) UpdateGroup(ctx context.Context, group domain.Group) error {
	if group.Author != common.UserIDFromContext(ctx) {
		return fmt.Errorf("%w: not allowed to modify others groups", domain.ErrNotAuthenticated)
	}

	if err := s.adapter.UpdateGroup(ctx, group); err != nil {
		return fmt.Errorf("s.adapter.UpdateGroup: %w", err)
	}

	return nil
}

func (s *Implementation) DeleteGroup(ctx context.Context, id domain.GroupID) error {
	group, err := s.adapter.GetGroup(ctx, id)
	if err != nil {
		return fmt.Errorf("s.adapter.GetGroup: %w", err)
	}

	if group.Author != common.UserIDFromContext(ctx) {
		return fmt.Errorf("%w: not allowed to modify others groups", domain.ErrNotAuthenticated)
	}

	if err := s.adapter.DeleteGroup(ctx, id); err != nil {
		return fmt.Errorf("s.adapter.DeleteGroup: %w", err)
	}

	return nil
}

func (s *Implementation) GetGroup(ctx context.Context, id domain.GroupID) (domain.Group, error) {
	group, err := s.adapter.GetGroup(ctx, id)
	if err != nil {
		return domain.Group{}, fmt.Errorf("s.adapter.GetGroup: %w", err)
	}

	return group, nil
}

func (s *Implementation) ListGroups(ctx context.Context) ([]domain.Group, error) {
	groups, err := s.adapter.ListGroups(ctx)
	if err != nil {
		return nil, fmt.Errorf("s.adapter.ListGroups: %w", err)
	}

	return groups, nil
}

func (i *Implementation) ListSubscribedGroups(ctx context.Context, id domain.UserID) ([]domain.GroupID, error) {
	groups, err := i.adapter.ListSubscribedGroups(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("s.adapter.ListSubscribedGroups: %w", err)
	}

	return groups, nil
}
