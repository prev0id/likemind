package auth

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"likemind/internal/common"
	"likemind/internal/domain"

	"github.com/gorilla/sessions"
)

const (
	sessionName = "app-session"
	userUUIDKey = "user-uuid"
)

func (i *implementation) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := i.ValidateSessionCookie(w, r)
		if err != nil {
			http.Redirect(w, r, domain.PathSignIn, http.StatusFound)
			return
		}

		ctx := common.ContextWithUserID(r.Context(), userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (i *implementation) ValidateSessionCookie(w http.ResponseWriter, r *http.Request) (int64, error) {
	session, err := i.cookieStore.Get(r, sessionName)
	if err != nil || session.IsNew {
		return 0, errors.New("no session cookie")
	}

	rawUUID, ok := session.Values[userUUIDKey]
	if !ok {
		return 0, errors.New("no uuid in session")
	}

	uuid, ok := rawUUID.(string)
	if !ok {
		i.invalidateSession(session, w, r)
		return 0, errors.New("invalid uuid type")
	}

	ctx := r.Context()

	creds, err := i.db.Get(ctx, domain.FieldUUID, uuid)
	if err != nil {
		i.invalidateSession(session, w, r)
		return 0, errors.New("invalid uuid in session")
	}

	return creds.UserID, nil
}

func (i *implementation) SetSessionCookie(uuid string, w http.ResponseWriter, r *http.Request) error {
	session, err := i.cookieStore.Get(r, sessionName)
	if err != nil {
		return fmt.Errorf("i.cookieStore.Get: %w", err)
	}

	session.Values[userUUIDKey] = uuid

	if err := session.Save(r, w); err != nil {
		return fmt.Errorf("session.Save: %w", err)
	}

	return nil
}

func (i *implementation) InvalidateSessionCookie(w http.ResponseWriter, r *http.Request) {
	session, err := i.cookieStore.Get(r, sessionName)
	if err != nil { // no cookie
		return
	}

	i.invalidateSession(session, w, r)
}

func (i *implementation) invalidateSession(session *sessions.Session, w http.ResponseWriter, r *http.Request) {
	session.Options.MaxAge = -1
	if err := session.Save(r, w); err != nil {
		log.Printf("session.Save: %s", err.Error())
	}
}
