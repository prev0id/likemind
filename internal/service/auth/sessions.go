package auth

import (
	"likemind/internal/common"
	"likemind/internal/domain"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/rs/zerolog/log"
)

const (
	sessionName = "SESSION"
	userIDKey   = "public-user-id"
)

func (i *implementation) ValidateSessionCookie(w http.ResponseWriter, r *http.Request) (int64, error) {
	session, err := i.cookieStore.Get(r, sessionName)
	if err != nil || session.IsNew {
		return 0, domain.ErrNotAuthenticated
	}

	rawID, ok := session.Values[userIDKey]
	if !ok {
		return 0, domain.ErrNotAuthenticated
	}

	id, ok := rawID.(string)
	if !ok {
		i.invalidateSession(session, w, r)
		return 0, domain.ErrNotAuthenticated
	}

	ctx := r.Context()

	creds, err := i.db.GetByID(ctx, id)
	if err != nil {
		i.invalidateSession(session, w, r)
		if common.ErrorIs(err, common.NotFoundErrorType) {
			err = domain.ErrNotAuthenticated
		}
		return 0, err
	}

	return creds.UserID, nil
}

func (i *implementation) SetSessionCookie(publicId string, w http.ResponseWriter, r *http.Request) error {
	session, err := i.cookieStore.Get(r, sessionName)
	if err != nil {
		return err
	}

	session.Values[userIDKey] = publicId

	if err := session.Save(r, w); err != nil {
		return err
	}

	return nil
}

func (i *implementation) InvalidateSessionCookie(w http.ResponseWriter, r *http.Request) {
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
