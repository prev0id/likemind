package auth

import (
	"fmt"
	"net/http"

	"likemind/internal/domain"

	"github.com/gorilla/sessions"
	"github.com/rs/zerolog/log"
)

const (
	sessionName = "app-session"
	userIDKey   = "public-user-id"
)

func (i *implementation) ValidateSessionCookie(w http.ResponseWriter, r *http.Request) (int64, error) {
	session, err := i.cookieStore.Get(r, sessionName)
	if err != nil || session.IsNew {
		return 0, domain.ErrNoSession
	}

	rawID, ok := session.Values[userIDKey]
	if !ok {
		return 0, domain.ErrInvalidSession
	}

	id, ok := rawID.(string)
	if !ok {
		i.invalidateSession(session, w, r)
		return 0, domain.ErrInvalidSession
	}

	ctx := r.Context()

	creds, err := i.db.GetByID(ctx, id)
	if err != nil {
		i.invalidateSession(session, w, r)
		return 0, domain.ErrInvalidSession
	}

	return creds.UserID, nil
}

func (i *implementation) SetSessionCookie(publicId string, w http.ResponseWriter, r *http.Request) error {
	session, err := i.cookieStore.Get(r, sessionName)
	if err != nil {
		return fmt.Errorf("i.cookieStore.Get: %w", err)
	}

	session.Values[userIDKey] = publicId

	if err := session.Save(r, w); err != nil {
		return fmt.Errorf("session.Save: %w", err)
	}

	return nil
}

func (i *implementation) IvalidateSessionCookie(w http.ResponseWriter, r *http.Request) {
	session, err := i.cookieStore.Get(r, sessionName)
	if err != nil {
		return
	}

	i.invalidateSession(session, w, r)
}

func (i *implementation) invalidateSession(session *sessions.Session, w http.ResponseWriter, r *http.Request) {
	session.Options.MaxAge = -1
	if err := session.Save(r, w); err != nil {
		log.Error().Err(err).Msg("session.Save")
	}
}
