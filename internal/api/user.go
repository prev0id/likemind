package api

import (
	"context"
	"net/url"

	"likemind/internal/common"
	"likemind/internal/domain"
	desc "likemind/internal/pkg/api"
)

func (s *Server) V1APILogoutPost(ctx context.Context) (desc.V1APILogoutPostRes, error) {
	userID := common.UserIDFromContext(ctx)

	cookie, err := s.session.InvalidateSession(ctx, userID)
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	return &desc.Redirect302{
		Location:  desc.NewOptURI(url.URL{Path: domain.PathPageSignIn}),
		SetCookie: desc.NewOptString(cookie.String()),
	}, nil
}

func (s *Server) V1APISignupPost(ctx context.Context, req *desc.SignUp) (desc.V1APISignupPostRes, error) {
	user := domain.User{
		Nickname:    req.GetUsername(),
		Name:        req.GetName(),
		Surname:     req.GetSurname(),
		Login:       domain.Email(req.GetEmail()),
		RawPassword: domain.Password(req.GetPassword()),
	}

	userID, err := s.profile.CreateUser(ctx, user)
	if common.ErrorIs(err, common.BadRequestErrorType) {
		return &desc.BadRequest{Data: common.ErrorMsg(err)}, nil
	}
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	cookie, err := s.session.CreateSessionCookie(ctx, userID)
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	return &desc.Redirect302{
		Location:  desc.NewOptURI(getProfilePage(user)),
		SetCookie: desc.NewOptString(cookie.String()),
	}, nil
}

func (s *Server) V1APISigninPost(ctx context.Context, req *desc.SignIn) (desc.V1APISigninPostRes, error) {
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
