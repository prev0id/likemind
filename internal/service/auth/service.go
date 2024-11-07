package auth

import (
	"context"
)

type Service interface{}

type db interface {
	ValidateSession(ctx context.Context, session_id uint64) error
}

type implementation struct {
	db db
}
