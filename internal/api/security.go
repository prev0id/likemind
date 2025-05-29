package api

import (
	"context"
	"fmt"
	"net/http"

	"likemind/internal/common"
	"likemind/internal/domain"
	desc "likemind/internal/pkg/api"
	"likemind/internal/service/session"

	"github.com/rs/zerolog/log"
)

type Security struct {
	sessions session.Service
}

var _ (desc.SecurityHandler) = (*Security)(nil)

func NewSecurityHandler(sessionService session.Service) *Security {
	return &Security{
		sessions: sessionService,
	}
}

func (s *Security) HandleSessionAuth(ctx context.Context, operationName desc.OperationName, t desc.SessionAuth) (context.Context, error) {
	userID, ok := common.UserIDFromContextWithCheck(ctx)
	if !ok {
		return ctx, fmt.Errorf("permission denied")
	}

	log.Warn().Msgf("permission granted, user=%d, operation=%q", userID, operationName)

	return ctx, nil
}

func (s *Security) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(domain.SessionName)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		ctx := r.Context()

		userID, err := s.sessions.ValidateSession(ctx, domain.SessionToken(cookie.Value))
		if err != nil {
			log.Err(err).Msg("s.sessions.ValidateSession")
			next.ServeHTTP(w, r)
			return
		}

		ctx = common.ContextWithUserID(ctx, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
