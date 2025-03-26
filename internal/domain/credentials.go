package domain

import "time"

const (
	SessionName = "SESSION"
)

type SessionToken string

type Session struct {
	UserID    UserID
	Token     SessionToken
	ExpiresAt time.Time
}
