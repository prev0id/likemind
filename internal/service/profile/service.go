package profile

import (
	"context"

	"likemind/internal/database/repo/contact_repo"
	profile_picture_repo "likemind/internal/database/repo/picture_repo"
	"likemind/internal/database/repo/user_repo"
	"likemind/internal/domain"
)

type Service interface {
	CreateUser(ctx context.Context, user domain.User) (int64, error)
	UpdateUser(ctx context.Context, user domain.User) error
	DeleteUser(ctx context.Context, id int64) error
	GetUser(ctx context.Context, id int64) (domain.User, error)
}

type implementation struct {
	userRepo    user_repo.DB
	pictureRepo profile_picture_repo.DB
	contactRepo contact_repo.DB
}
