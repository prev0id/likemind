package group

import (
	"context"
	"likemind/internal/database/adapter/group_adapter"
	"likemind/internal/domain"
)

type Service interface {
	CreateGroup(ctx context.Context, group domain.Group) (domain.GroupID, error)
	UpdateGroup(ctx context.Context, group domain.Group) error
	DeleteGroup(ctx context.Context, id domain.GroupID) error
	GetGroup(ctx context.Context, id domain.GroupID) (domain.Group, error)
	ListGroups(ctx context.Context) ([]domain.Group, error)
	ListSubscribedGroups(ctx context.Context, id domain.UserID) ([]domain.GroupID, error)

	CreatePost(ctx context.Context, post domain.Post) (domain.PostID, error)
	UpdatePost(ctx context.Context, post domain.Post) error
	DeletePost(ctx context.Context, id domain.PostID) error
	GetPosts(ctx context.Context, id domain.GroupID) ([]domain.Post, error)

	CreateComment(ctx context.Context, comment domain.Comment) (domain.CommentID, error)
	UpdateComment(ctx context.Context, comment domain.Comment) error
	DeleteComment(ctx context.Context, id domain.CommentID) error
	GetComments(ctx context.Context, id domain.PostID) ([]domain.Comment, error)
}

type Implementation struct {
	adapter group_adapter.Adapter
}

var _ Service = (*Implementation)(nil)

func New(adapter group_adapter.Adapter) *Implementation {
	return &Implementation{
		adapter: adapter,
	}
}
