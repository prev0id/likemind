package auth

import "net/http"

const sessionCookie = "session_token"

func (i *implementation) Middlware(next *http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := r.Cookie(sessionCookie)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		i.db.ValidateUser(ctx context.Context, token string)
	})
}
