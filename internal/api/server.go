package api

import (
	"context"
	"net/url"

	"likemind/internal/common"
	"likemind/internal/domain"
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

func (s *Server) V1APISigninPost(ctx context.Context, req *desc.V1APISigninPostReq) (desc.V1APISigninPostRes, error) {
	if req == nil {
		return nil, domain.ErrNilRequest
	}

	var (
		email    = domain.Email(req.Email)
		password = domain.Password(req.Password)
	)

	user, err := s.profile.SignIn(ctx, email, password)
	if common.ErrorIs(err, common.NotAuthenticatedErrorType) {
		return &desc.NotAuthorized{Data: common.ErrorMsg(err)}, nil
	}
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	cookie, err := s.session.CreateSessionCookie(ctx, user.ID)
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	return &desc.Redirect302{
		Location:  desc.NewOptURI(getProfilePage(user)),
		SetCookie: desc.NewOptString(cookie.String()),
	}, nil
}

func getProfilePage(user domain.User) url.URL {
	params := map[string]string{
		domain.PathParamUsername: user.Nickname,
	}

	path := common.FillPath(domain.PathPageUser, params)

	return url.URL{
		Path: path,
	}
}
