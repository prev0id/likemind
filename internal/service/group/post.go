package group

import (
	"context"
	"fmt"

	"likemind/internal/common"
	"likemind/internal/domain"
)

func (s *Implementation) CreatePost(ctx context.Context, post domain.Post) (domain.PostID, error) {
	post.Author = common.UserIDFromContext(ctx)

	id, err := s.adapter.CreatePost(ctx, post)
	if err != nil {
		return 0, fmt.Errorf("s.adapter.CreatePost: %w", err)
	}

	return id, nil
}

func (s *Implementation) UpdatePost(ctx context.Context, post domain.Post) error {
	if post.Author != common.UserIDFromContext(ctx) {
		return fmt.Errorf("%w: not allowed to modify others posts", domain.ErrNotAuthenticated)
	}

	if err := s.adapter.UpdatePost(ctx, post); err != nil {
		return fmt.Errorf("s.adapter.UpdatePost: %w", err)
	}

	return nil
}

func (s *Implementation) DeletePost(ctx context.Context, id domain.PostID) error {
	post, err := s.adapter.GetPostByID(ctx, id)
	if err != nil {
		return fmt.Errorf("s.adapter.GetPostByID: %w", err)
	}

	if post.Author != common.UserIDFromContext(ctx) {
		return fmt.Errorf("%w: not allowed to modify others posts", domain.ErrNotAuthenticated)
	}

	if err := s.adapter.DeletePost(ctx, id); err != nil {
		return fmt.Errorf("s.adapter.DeletePost: %w", err)
	}

	return nil
}

func (s *Implementation) GetPosts(ctx context.Context, groupID domain.GroupID) ([]domain.Post, error) {
	posts, err := s.adapter.GetGroupPosts(ctx, groupID)
	if err != nil {
		return nil, fmt.Errorf("s.adapter.GetGroupPosts: %w", err)
	}

	return posts, nil
}
