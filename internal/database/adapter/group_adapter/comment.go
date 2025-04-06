package group_adapter

import (
	"context"
	"fmt"

	"likemind/internal/common"
	"likemind/internal/domain"
)

func (i *Implementation) CreateComment(ctx context.Context, comment domain.Comment) (domain.CommentID, error) {
	id, err := i.comment.Create(ctx, commentDomainToModel(comment))
	if err != nil {
		return 0, fmt.Errorf("i.comment.Create: %w", err)
	}

	return domain.CommentID(id), nil
}

func (i *Implementation) UpdateComment(ctx context.Context, comment domain.Comment) error {
	if err := i.comment.Update(ctx, commentDomainToModel(comment)); err != nil {
		return fmt.Errorf("i.comment.Update: %w", err)
	}

	return nil
}

func (i *Implementation) DeleteComment(ctx context.Context, id domain.CommentID) error {
	if err := i.comment.Delete(ctx, int64(id)); err != nil {
		return fmt.Errorf("i.comment.Delete: %w", err)
	}

	return nil
}

func (i *Implementation) GetPostComments(ctx context.Context, postID domain.PostID) ([]domain.Comment, error) {
	comments, err := i.comment.ListByPost(ctx, int64(postID))
	if err != nil {
		return nil, fmt.Errorf("i.comment.ListByPost: %w", err)
	}

	return common.Convert(comments, commentModelToDomain), nil
}

func (i *Implementation) GetCommentByID(ctx context.Context, id domain.CommentID) (domain.Comment, error) {
	comment, err := i.comment.GetByID(ctx, int64(id))
	if err != nil {
		return domain.Comment{}, fmt.Errorf("i.comment.GetByID: %w", err)
	}

	return commentModelToDomain(comment), nil
}
