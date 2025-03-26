package session_adapter

import (
	"time"

	"likemind/internal/database/model"
	"likemind/internal/domain"
)

func modelSessionToDomain(s model.Session) domain.Session {
	return domain.Session{
		UserID:    domain.UserID(s.UserID),
		Token:     domain.SessionToken(s.Token),
		ExpiresAt: s.ExpiresAt,
	}
}

func domainSessionToModel(s domain.Session) model.Session {
	return model.Session{
		UserID:    int64(s.UserID),
		Token:     string(s.Token),
		CreatedAt: time.Now(),
		ExpiresAt: s.ExpiresAt,
	}
}
