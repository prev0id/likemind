package api

import (
	"context"
	"fmt"
	"likemind/internal/domain"
	desc "likemind/internal/pkg/api"
	"likemind/internal/service/group"
	"likemind/internal/service/image"
	"likemind/internal/service/interests"
	"likemind/internal/service/profile"
	"likemind/internal/service/session"
	"likemind/website/view"
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

func (s *Server) getProfile(ctx context.Context, userID domain.UserID) (*view.Profile, error) {
	profile, err := s.profile.GetUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("s.profile.GetUser: %w", err)
	}

	pictures, err := s.image.GetProfileImages(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("s.image.GetProfileImages: %w", err)
	}

	contacts, err := s.profile.GetContacts(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("s.profile.GetContacts: %w", err)
	}

	interests, err := s.interests.GetUserInterests(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("s.interests.GetUserInterests: %w", err)
	}

	return profileFromDomainToView(userID, profile, contacts, pictures, interests), nil
}

func (s *Server) getGroup(ctx context.Context, groupID domain.GroupID) (*view.Group, error) {
	group, err := s.group.GetGroup(ctx, groupID)
	if err != nil {
		return nil, fmt.Errorf("s.group.GetGroup: %w", err)
	}

	posts, err := s.group.GetPosts(ctx, groupID)
	if err != nil {
		return nil, fmt.Errorf("s.group.GetPosts: %w", err)
	}

	interests, err := s.interests.GetGroupInterests(ctx, groupID)
	if err != nil {
		return nil, fmt.Errorf("s.interests.GetGroupInterests: %w", err)
	}

	return s.groupDomainToView(ctx, group, posts, interests)
}
