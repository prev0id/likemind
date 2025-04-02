package api

import (
	"context"
	"net/url"

	"likemind/internal/common"
	"likemind/internal/domain"
	desc "likemind/internal/pkg/api"
)

func (s *Server) V1APIProfilePost(ctx context.Context, req *desc.ProfileCreate) (desc.V1APIProfilePostRes, error) {
	user := domain.User{
		Nickname:    req.GetUsername(),
		Name:        req.GetName(),
		Surname:     req.GetSurname(),
		Login:       domain.Email(req.GetEmail()),
		RawPassword: domain.Password(req.GetPassword()),
		DateOfBirth: req.GetDateOfBirth(),
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
		HxRedirect: desc.NewOptURI(getProfilePage(user)),
		SetCookie:  desc.NewOptString(cookie.String()),
	}, nil
}

func (s *Server) V1APIProfileDelete(ctx context.Context) (desc.V1APIProfileDeleteRes, error) {
	userID := common.UserIDFromContext(ctx)

	cookie, err := s.session.InvalidateSession(ctx, userID)
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	if err := s.profile.DeleteProfile(ctx, userID); err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	return &desc.Redirect302{
		HxRedirect: desc.NewOptURI(url.URL{Path: domain.PathPageSignUp}),
		SetCookie:  desc.NewOptString(cookie.String()),
	}, nil
}

func (s *Server) V1APIProfilePut(ctx context.Context, req *desc.ProfileUpdate) (desc.V1APIProfilePutRes, error) {
	userID := common.UserIDFromContext(ctx)

	user, err := s.profile.GetUser(ctx, userID)
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	user = updateUserFields(user, req)

	if err := s.profile.UpdateUser(ctx, user); err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	return &desc.HTMLResponse{}, nil
}

func getProfilePage(user domain.User) url.URL {
	params := map[string]string{
		domain.PathParamUsername: user.Nickname,
	}

	path := common.FillPath(domain.PathPageProfile, params)

	return url.URL{
		Path: path,
	}
}

func updateUserFields(user domain.User, req *desc.ProfileUpdate) domain.User {
	if email := req.GetEmail(); email.IsSet() {
		user.Login = domain.Email(email.Value)
	}
	if name := req.GetName(); name.IsSet() {
		user.Name = name.Value
	}
	if surname := req.GetSurname(); surname.IsSet() {
		user.Surname = surname.Value
	}
	if username := req.GetUsername(); username.IsSet() {
		user.Nickname = username.Value
	}
	if dateOfBirth := req.GetDateOfBirth(); dateOfBirth.IsSet() {
		user.DateOfBirth = dateOfBirth.Value
	}
	if about := req.GetAbout(); about.IsSet() {
		user.About = about.Value
	}
	if location := req.GetLocation(); location.IsSet() {
		user.Location = location.Value
	}
	return user
}
