package group

import (
	"context"
	"fmt"

	"likemind/internal/common"
	"likemind/internal/domain"
)

func (s *Implementation) CreateComment(ctx context.Context, comment domain.Comment) (domain.CommentID, error) {
	comment.Author = common.UserIDFromContext(ctx)

	id, err := s.adapter.CreateComment(ctx, comment)
	if err != nil {
		return 0, fmt.Errorf("s.adapter.CreateComment: %w", err)
	}

	return id, nil
}

func (s *Implementation) UpdateComment(ctx context.Context, comment domain.Comment) error {
	if comment.Author != common.UserIDFromContext(ctx) {
		return fmt.Errorf("%w: not allowed to modify others comments", domain.ErrNotAuthenticated)
	}

	if err := s.adapter.UpdateComment(ctx, comment); err != nil {
		return fmt.Errorf("s.adapter.UpdateComment: %w", err)
	}

	return nil
}

func (s *Implementation) DeleteComment(ctx context.Context, id domain.CommentID) error {
	comment, err := s.adapter.GetCommentByID(ctx, id)
	if err != nil {
		return fmt.Errorf("s.adapter.GetCommentByID: %w", err)
	}

	if comment.Author != common.UserIDFromContext(ctx) {
		return fmt.Errorf("%w: not allowed to modify others comments", domain.ErrNotAuthenticated)
	}

	if err := s.adapter.DeleteComment(ctx, id); err != nil {
		return fmt.Errorf("s.adapter.DeleteComment: %w", err)
	}

	return nil
}

func (s *Implementation) GetComments(ctx context.Context, postID domain.PostID) ([]domain.Comment, error) {
	comments, err := s.adapter.GetPostComments(ctx, postID)
	if err != nil {
		return nil, fmt.Errorf("s.adapter.GetPostComments: %w", err)
	}

	return comments, nil
}
