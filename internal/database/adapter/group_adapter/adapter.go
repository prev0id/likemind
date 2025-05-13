package group_adapter

import (
	"context"
	"likemind/internal/database/repo/comment_repo"
	"likemind/internal/database/repo/group_repo"
	"likemind/internal/database/repo/post_repo"
	"likemind/internal/domain"
)

type Adapter interface {
	CreateGroup(ctx context.Context, group domain.Group) (domain.GroupID, error)
	UpdateGroup(ctx context.Context, group domain.Group) error
	DeleteGroup(ctx context.Context, id domain.GroupID) error
	GetGroup(ctx context.Context, id domain.GroupID) (domain.Group, error)
	ListGroups(ctx context.Context) ([]domain.Group, error)
	ListSubscribedGroups(ctx context.Context, id domain.UserID) ([]domain.GroupID, error)

	CreatePost(ctx context.Context, post domain.Post) (domain.PostID, error)
	UpdatePost(ctx context.Context, post domain.Post) error
	DeletePost(ctx context.Context, id domain.PostID) error
	GetGroupPosts(ctx context.Context, id domain.GroupID) ([]domain.Post, error)
	GetPostByID(ctx context.Context, id domain.PostID) (domain.Post, error)

	CreateComment(ctx context.Context, comment domain.Comment) (domain.CommentID, error)
	UpdateComment(ctx context.Context, comment domain.Comment) error
	DeleteComment(ctx context.Context, id domain.CommentID) error
	GetPostComments(ctx context.Context, id domain.PostID) ([]domain.Comment, error)
	GetCommentByID(ctx context.Context, id domain.CommentID) (domain.Comment, error)
}

type Implementation struct {
	group   group_repo.DB
	post    post_repo.DB
	comment comment_repo.DB
}

var _ (Adapter) = (*Implementation)(nil)

func New(
	group group_repo.DB,
	post post_repo.DB,
	comment comment_repo.DB,
) *Implementation {
	return &Implementation{
		group:   group,
		post:    post,
		comment: comment,
	}
}
