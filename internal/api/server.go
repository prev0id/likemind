package api

import (
	desc "likemind/internal/pkg/api"
	"likemind/internal/service/group"
	"likemind/internal/service/image"
	"likemind/internal/service/interests"
	"likemind/internal/service/profile"
	"likemind/internal/service/session"
)

type Server struct {
	session   session.Service
	profile   profile.Service
	group     group.Service
	image     image.Service
	interests interests.Service
}

var _ (desc.Handler) = (*Server)(nil)

func NewServer(
	session session.Service,
	profile profile.Service,
	image image.Service,
	interests interests.Service,
) *Server {
	return &Server{
		session:   session,
		profile:   profile,
		image:     image,
		interests: interests,
	}
}
