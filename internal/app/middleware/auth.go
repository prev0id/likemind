package middleware

import (
	"net/http"

	"likemind/internal/common"
	"likemind/internal/domain"
	"likemind/internal/service/auth"
)

type Middleware func(next http.Handler) http.Handler

func NewAuthMiddleware(svc auth.Service) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID, err := svc.ValidateSessionCookie(w, r)
			if err != nil {
				http.Redirect(w, r, domain.PathSignIn, http.StatusFound)
				return
			}

			ctx := common.ContextWithUserID(r.Context(), userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
