package auth

import (
	"context"
	"fmt"
	"net/http"

	"likemind/internal/config"
	"likemind/internal/database"
	"likemind/internal/database/adapter/credentials_adapter"
	"likemind/internal/domain"

	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/stdlib"
)

type Service interface {
	SetSessionCookie(uuid string, w http.ResponseWriter, r *http.Request) error
	InvalidateSessionCookie(w http.ResponseWriter, r *http.Request)
	ValidateSessionCookie(w http.ResponseWriter, r *http.Request) (int64, error)
	NewUserCredentials(ctx context.Context, userID int64, login domain.Email, password domain.Password) (string, error)
	Signin(ctx context.Context, login domain.Email, password domain.Password) (domain.Credentials, error)
	Authenticate(ctx context.Context, credsID string) (int64, error)
	Close()
}

type implementation struct {
	db          credentials_adapter.Adapter
	cookieStore *pgstore.PGStore
}

func New(db credentials_adapter.Adapter, cfg config.Auth) (Service, error) {
	cookieStore, err := pgstore.NewPGStoreFromPool(stdlib.OpenDBFromPool(database.DB), []byte(cfg.SessionKey))
	if err != nil {
		return nil, fmt.Errorf("pgstore.NewPGStoreFromPool: %w", err)
	}

	cookieStore.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   cfg.CookieMaxAge,
		HttpOnly: cfg.UseHTTPOnly,
		Secure:   true,
	}

	return &implementation{
		db:          db,
		cookieStore: cookieStore,
	}, nil
}

func (i *implementation) Close() {
	i.cookieStore.Close()
}
