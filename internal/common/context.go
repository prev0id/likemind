package common

import (
	"context"

	"likemind/internal/domain"
)

type userIDType struct{}

var userIDKey userIDType

func ContextWithUserID(ctx context.Context, id domain.UserID) context.Context {
	return context.WithValue(ctx, userIDKey, id)
}

func UserIDFromContext(ctx context.Context) domain.UserID {
	userID, ok := ctx.Value(userIDKey).(domain.UserID)
	if !ok {
		return 0
	}
	return userID
}

func UserIDFromContextWithCheck(ctx context.Context) (domain.UserID, bool) {
	userID, ok := ctx.Value(userIDKey).(domain.UserID)
	if !ok {
		return 0, false
	}
	return userID, true
}
