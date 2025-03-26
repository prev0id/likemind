package api

import (
	"context"
	"fmt"

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
	token := domain.SessionToken(t.GetAPIKey())

	userID, err := s.sessions.ValidateSession(ctx, token)
	if err != nil {
		log.Warn().Msgf("permission denied, token=%q, reason=%q", token, err.Error())
		return ctx, fmt.Errorf("s.sessions.ValidateSession: %w", err)
	}

	log.Warn().Msgf("permission granted, user=%d, operation=%q", userID, operationName)
	return common.ContextWithUserID(ctx, userID), nil
}
