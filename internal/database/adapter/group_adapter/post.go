package group_adapter

import (
	"context"
	"fmt"

	"likemind/internal/database"
	"likemind/internal/domain"
)

func (i *Implementation) CreatePost(ctx context.Context, post domain.Post) (domain.PostID, error) {
	id, err := i.post.Create(ctx, postDomainToModel(post))
	if err != nil {
		return 0, fmt.Errorf("i.post.Create: %w", err)
	}
	return domain.PostID(id), nil
}

func (i *Implementation) UpdatePost(ctx context.Context, post domain.Post) error {
	if err := i.post.Update(ctx, postDomainToModel(post)); err != nil {
		return fmt.Errorf("i.post.Update: %w", err)
	}
	return nil
}

func (i *Implementation) DeletePost(ctx context.Context, id domain.PostID) error {
	err := database.InTransaction(ctx, func(ctx context.Context) error {
		comments, err := i.comment.ListByPost(ctx, int64(id))
		if err != nil {
			return fmt.Errorf("i.post.ListByPost: %w", err)
		}

		for _, comment := range comments {
			if err := i.DeleteComment(ctx, domain.CommentID(comment.ID)); err != nil {
				return fmt.Errorf("i.DeleteComment: %w", err)
			}
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("database.InTransaction: %w", err)
	}

	return nil
}

func (i *Implementation) GetPostByID(ctx context.Context, id domain.PostID) (domain.Post, error) {
	post, err := i.post.GetByID(ctx, int64(id))
	if err != nil {
		return domain.Post{}, fmt.Errorf("i.post.GetByID: %w", err)
	}

	return postModelToDomain(post), err
}
