package group_adapter

import (
	"context"
	"fmt"

	"likemind/internal/common"
	"likemind/internal/database"
	"likemind/internal/database/model"
	"likemind/internal/domain"
)

func (i *Implementation) CreateGroup(ctx context.Context, group domain.Group) (domain.GroupID, error) {
	id, err := i.group.Create(ctx, groupDomainToModel(group))
	if err != nil {
		return 0, fmt.Errorf("i.group.Create: %w", err)
	}
	return domain.GroupID(id), nil
}

func (i *Implementation) UpdateGroup(ctx context.Context, group domain.Group) error {
	if err := i.group.Update(ctx, groupDomainToModel(group)); err != nil {
		return fmt.Errorf("i.group.Update: %w", err)
	}

	return nil
}

func (i *Implementation) DeleteGroup(ctx context.Context, id domain.GroupID) error {
	err := database.InTransaction(ctx, func(ctx context.Context) error {
		posts, err := i.post.ListByGroupID(ctx, int64(id))
		if err != nil {
			return fmt.Errorf("i.post.ListByGroupID: %w", err)
		}

		for _, post := range posts {
			if err := i.DeletePost(ctx, domain.PostID(post.ID)); err != nil {
				return fmt.Errorf("i.DeletePost: %w", err)
			}
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("database.InTransaction: %w", err)
	}

	return nil
}

func (i *Implementation) GetGroup(ctx context.Context, id domain.GroupID) (domain.Group, error) {
	m, err := i.group.GetByID(ctx, int64(id))
	if err != nil {
		return domain.Group{}, fmt.Errorf("i.group.GetByID: %w", err)
	}

	return groupModelToDomain(m), nil
}

func (i *Implementation) ListGroups(ctx context.Context) ([]domain.Group, error) {
	groups, err := i.group.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("i.group.List: %w", err)
	}

	return common.Convert(groups, groupModelToDomain), nil
}

func (i *Implementation) GetGroupPosts(ctx context.Context, groupID domain.GroupID) ([]domain.Post, error) {
	posts, err := i.post.ListByGroupID(ctx, int64(groupID))
	if err != nil {
		return nil, fmt.Errorf("i.post.ListByGroupID: %w", err)
	}

	return common.Convert(posts, postModelToDomain), nil
}

func (i *Implementation) ListSubscribedGroups(ctx context.Context, id domain.UserID) ([]domain.GroupID, error) {
	subscriptions, err := i.group.ListUserSubscriptions(ctx, int64(id))
	if err != nil {
		return nil, fmt.Errorf("i.group.ListUserSubscriptions: %w", err)
	}

	result := make([]domain.GroupID, 0, len(subscriptions))
	for _, subscription := range subscriptions {
		result = append(result, domain.GroupID(subscription.GroupID))
	}

	return result, nil
}

func (i *Implementation) Subscribe(ctx context.Context, userID domain.UserID, groupID domain.GroupID) error {
	return i.group.AddUserSubscription(ctx, model.UserSubscription{
		UserID:  int64(userID),
		GroupID: int64(groupID),
	})
}

func (i *Implementation) Unsubscribe(ctx context.Context, userID domain.UserID, groupID domain.GroupID) error {
	return i.group.DeleteUserSubscription(ctx, model.UserSubscription{
		UserID:  int64(userID),
		GroupID: int64(groupID),
	})
}
