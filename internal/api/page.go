package api

import (
	"context"
	"net/http"

	"likemind/internal/common"
	desc "likemind/internal/pkg/api"
	error_page "likemind/website/page/error"
	group_page "likemind/website/page/group"
	profile_page "likemind/website/page/profile"
	signin_page "likemind/website/page/signin"
	signup_page "likemind/website/page/signup"
	user_search_page "likemind/website/page/user_search"
)

func (s *Server) V1PageGroupGroupNameGet(ctx context.Context, params desc.V1PageGroupGroupNameGetParams) (desc.V1PageGroupGroupNameGetRes, error) {
	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, group_page.Page()),
	}, nil
}

func (s *Server) V1PageProfileUsernameGet(ctx context.Context, params desc.V1PageProfileUsernameGetParams) (desc.V1PageProfileUsernameGetRes, error) {
	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, profile_page.Page()),
	}, nil
}

func (s *Server) V1PageSearchGet(ctx context.Context) (desc.V1PageSearchGetRes, error) {
	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, user_search_page.Page()),
	}, nil
}

func (s *Server) V1PageSigninGet(ctx context.Context) (desc.V1PageSigninGetRes, error) {
	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, signin_page.Page()),
	}, nil
}

func (s *Server) V1PageSignupGet(ctx context.Context) (desc.V1PageSignupGetRes, error) {
	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, signup_page.Page()),
	}, nil
}

func (s *Server) NotFound(w http.ResponseWriter, r *http.Request) {
	page := error_page.Page(error_page.State{Code: http.StatusNotFound})

	w.WriteHeader(http.StatusNotFound)
	page.Render(r.Context(), w)
}
