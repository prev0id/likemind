package api

import (
	"context"
	"errors"
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
		HxRedirect: getProfilePage(),
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

func (s *Server) V1APIProfileEmailPut(ctx context.Context, req *desc.EmailUpdate) (desc.V1APIProfileEmailPutRes, error) {
	userID := common.UserIDFromContext(ctx)

	err := s.profile.UpdateEmail(
		ctx,
		userID,
		domain.Email(req.Email),
		domain.Email(req.NewEmail),
		domain.Password(req.Password),
	)
	if err != nil {
		return &desc.BadRequest{Data: common.ErrorMsg(err)}, nil
	}

	return &desc.Redirect302{
		HxRedirect: getProfilePage(),
	}, nil
}

func (s *Server) V1APIProfileImageImageIDGet(ctx context.Context, params desc.V1APIProfileImageImageIDGetParams) (desc.V1APIProfileImageImageIDGetRes, error) {
	userID := common.UserIDFromContext(ctx)

	img, err := s.image.GetImage(ctx, domain.PictureID(params.ImageID), userID)
	if err != nil {
		return &desc.NotFound{Data: common.ErrorMsg(err)}, nil
	}

	return &desc.ImageImageJpeg{Data: img}, nil
}

func (s *Server) V1APIProfileImagePost(ctx context.Context, req *desc.V1APIProfileImagePostReq, params desc.V1APIProfileImagePostParams) (desc.V1APIProfileImagePostRes, error) {
	file, ok := req.GetImage().Get()
	if !ok {
		return &desc.InternalError{Data: common.ErrorMsg(errors.New("fuuuuuck"))}, nil
	}

	if err := s.image.UploadImage(ctx, file); err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	return &desc.Redirect302{
		HxRedirect: getProfilePage(),
	}, nil
}

func (s *Server) V1APIProfilePasswordPut(ctx context.Context, req *desc.PasswordUpdate) (desc.V1APIProfilePasswordPutRes, error) {
	userID := common.UserIDFromContext(ctx)

	err := s.profile.UpdatePassword(
		ctx,
		userID,
		domain.Email(req.Email),
		domain.Password(req.Password),
		domain.Password(req.NewPassword),
	)
	if err != nil {
		return &desc.BadRequest{Data: common.ErrorMsg(err)}, nil
	}

	return &desc.Redirect302{
		HxRedirect: getProfilePage(),
	}, nil
}

func getProfilePage() desc.OptURI {
	return desc.NewOptURI(url.URL{
		Path: domain.PathPageOwnProfile,
	})
}

func updateUserFields(user domain.User, req *desc.ProfileUpdate) domain.User {
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

// func getImageData(req desc.V1APIProfileImagePostReq) (io.Reader, bool) {
// 	switch request := req.(type) {
// 	case *desc.V1APIProfileImagePostReqImageJpeg:
// 		return request.Data, true
// 	case *desc.V1APIProfileImagePostReqImagePNG:
// 		return request.Data, true
// 	default:
// 		return nil, false
// 	}
// }
