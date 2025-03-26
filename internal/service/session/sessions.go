package session

import (
	"context"
	"crypto/rand"
	"fmt"
	"net/http"
	"time"

	"likemind/internal/domain"
)

const (
	noToken = ""
	expire  = -1
)

func (i *implementation) ValidateSession(ctx context.Context, token domain.SessionToken) (domain.UserID, error) {
	session, err := i.db.GetByToken(ctx, token)
	if err != nil {
		return 0, fmt.Errorf("i.db.GetByToken: %w", err)
	}

	return session.UserID, nil
}

func (i *implementation) InvalidateSession(ctx context.Context, id domain.UserID) (*http.Cookie, error) {
	if err := i.db.Invalidate(ctx, id); err != nil {
		return nil, fmt.Errorf("i.db.Invalidate: %w", err)
	}

	return createCookie(noToken, expire), nil
}

func (i *implementation) CreateSessionCookie(ctx context.Context, id domain.UserID) (*http.Cookie, error) {
	token := rand.Text()

	session := domain.Session{
		UserID:    id,
		Token:     domain.SessionToken(token),
		ExpiresAt: time.Now().Add(i.exparation),
	}

	if err := i.db.Create(ctx, session); err != nil {
		return nil, fmt.Errorf("i.db.Create: %w", err)
	}

	return createCookie(token, int(i.exparation.Seconds())), nil
}

func createCookie(token string, maxAge int) *http.Cookie {
	return &http.Cookie{
		Name:     domain.SessionName,
		Value:    token,
		Path:     "/",
		MaxAge:   maxAge,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
}
