package api

import (
	"context"
	"net/http"

	"likemind/internal/common"
	desc "likemind/internal/pkg/api"
	"likemind/website/page"
	error_page "likemind/website/page/error"
	group_page "likemind/website/page/group"
	user_search_page "likemind/website/page/user_search"
)

func (s *Server) V1PageGroupGroupIDGet(ctx context.Context, params desc.V1PageGroupGroupIDGetParams) (desc.V1PageGroupGroupIDGetRes, error) {
	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, group_page.Page(group_page.State)),
	}, nil
}

func (s *Server) V1PageProfileUsernameGet(ctx context.Context, params desc.V1PageProfileUsernameGetParams) (desc.V1PageProfileUsernameGetRes, error) {
	userID := common.UserIDFromContext(ctx)

	profile, err := s.profile.GetUser(ctx, userID)
	if err != nil {
		if common.ErrorIs(err, common.NotFoundErrorType) {
			return &desc.NotFound{Data: common.ErrorMsg(err)}, nil
		}
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	pictures, err := s.profile.GetProfilePictures(ctx, userID)
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	contacts, err := s.profile.GetContacts(ctx, userID)
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	state := profileFromDomainToView(ctx, profile, contacts, pictures)
	pageComponent := page.Profile(state)

	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, pageComponent),
	}, nil
}

func (s *Server) V1PageSearchGet(ctx context.Context) (desc.V1PageSearchGetRes, error) {
	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, user_search_page.Page()),
	}, nil
}

func (s *Server) V1PageSigninGet(ctx context.Context, params desc.V1PageSigninGetParams) (desc.V1PageSigninGetRes, error) {
	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, page.SignIn()),
	}, nil
}

func (s *Server) V1PageSignupGet(ctx context.Context, params desc.V1PageSignupGetParams) (desc.V1PageSignupGetRes, error) {
	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, page.SignUp()),
	}, nil
}

func (s *Server) NotFound(w http.ResponseWriter, r *http.Request) {
	page := error_page.Page(error_page.State{Code: http.StatusNotFound})

	w.WriteHeader(http.StatusNotFound)
	page.Render(r.Context(), w)
}
