package session

import (
	"context"
	"net/http"
	"time"

	"likemind/internal/config"
	"likemind/internal/database/adapter/session_adapter"
	"likemind/internal/domain"
)

type Service interface {
	ValidateSession(ctx context.Context, token domain.SessionToken) (domain.UserID, error)
	InvalidateSession(ctx context.Context, id domain.UserID) (*http.Cookie, error)
	CreateSessionCookie(ctx context.Context, id domain.UserID) (*http.Cookie, error)
}

type implementation struct {
	db         session_adapter.Adapter
	exparation time.Duration
}

func New(db session_adapter.Adapter, cfg config.Auth) Service {
	return &implementation{
		db:         db,
		exparation: cfg.Exparation,
	}
}
