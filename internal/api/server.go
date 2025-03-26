package api

import (
	desc "likemind/internal/pkg/api"
	"likemind/internal/service/profile"
	"likemind/internal/service/session"
)

type Server struct {
	session session.Service
	profile profile.Service
}

var _ (desc.Handler) = (*Server)(nil)

func NewServer(session session.Service, profile profile.Service) *Server {
	return &Server{
		session: session,
		profile: profile,
	}
}
