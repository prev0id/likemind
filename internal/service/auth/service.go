package auth

import (
	"context"
	"net/http"
)

type Service interface {
	Middlware(next *http.Handler) http.Handler
}

type db interface {
	ValidateSession(ctx context.Context, session_id uint64) error
}

type implementation struct {
	db db
}
