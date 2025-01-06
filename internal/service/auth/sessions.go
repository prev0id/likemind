package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"likemind/internal/common"
	"likemind/internal/domain"

	"github.com/gorilla/sessions"
	"github.com/samber/lo"
)

const (
	sessionName = "app-session"
	userUUIDKey = "user-uuid"
)

func (i *implementation) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := i.cookieStore.Get(r, sessionName)
		if err != nil || session.IsNew {
			http.Redirect(w, r, domain.PathSignIn, http.StatusFound)
			return
		}

		rawUUID, ok := session.Values[userUUIDKey]
		if !ok {
			http.Redirect(w, r, domain.PathSignIn, http.StatusFound)
			return
		}

		uuid, ok := rawUUID.(string)
		if !ok {
			i.invalidateSession(session, w, r)
			http.Redirect(w, r, domain.PathSignIn, http.StatusFound)
			return
		}

		ctx := r.Context()

		creds, err := i.credentialByUUID(ctx, uuid)
		if err != nil {
			i.invalidateSession(session, w, r)
			http.Redirect(w, r, domain.PathSignIn, http.StatusFound)
			return
		}

		ctx = common.ContextWithUserID(ctx, creds.UserID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
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

func (i *implementation) invalidateSession(session *sessions.Session, w http.ResponseWriter, r *http.Request) {
	session.Options.MaxAge = -1
	if err := session.Save(r, w); err != nil {
		log.Printf("session.Save: %s", err.Error())
	}
}

func (i *implementation) credentialByUUID(ctx context.Context, uuid string) (domain.Credential, error) {
	creds, err := i.db.List(ctx)
	if err != nil {
		return domain.Credential{}, fmt.Errorf("i.db.List: %w", err)
	}

	cred, found := lo.Find(creds, func(cred domain.Credential) bool {
		return cred.UUID == uuid
	})

	if !found {
		return domain.Credential{}, fmt.Errorf("user with uuid '%s' not found", uuid)
	}

	return cred, nil
}
