package auth

import (
	"context"
	"fmt"
	"net/http"

	"likemind/internal/config"
	"likemind/internal/domain"

	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"golang.org/x/crypto/bcrypt"
)

const (
	saltPattern = "%d$%s"
	cost        = bcrypt.DefaultCost

	passwordMaxLen = 50
	passwordMinLen = 5

	numbers      = "0123456789"
	specialChars = "$@!%*#?&"
)

type Service interface {
	Middleware(http.Handler) http.Handler
	SetSessionCookie(uuid string, w http.ResponseWriter, r *http.Request) error
	SetCredentials(ctx context.Context, userID int64, login, password string) (string, error)
	ValidateCredentials(ctx context.Context, login, password string) (string, error)
	Close()
}

type implementation struct {
	db          domain.DataProvider[domain.Credential]
	cookieStore *pgstore.PGStore
}

func New(provider domain.DataProvider[domain.Credential], dbConn *pgxpool.Pool, cfg config.Auth) (*implementation, error) {
	cookieStore, err := pgstore.NewPGStoreFromPool(stdlib.OpenDBFromPool(dbConn), []byte(cfg.SessionKey))
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
		db:          provider,
		cookieStore: cookieStore,
	}, nil
}

func (i *implementation) Close() {
	i.cookieStore.Close()
}
