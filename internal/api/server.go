package api

import (
	"context"
	"likemind/internal/common"
	desc "likemind/internal/pkg/api"
	"likemind/internal/service/auth"
	"likemind/internal/service/profile"
	group_page "likemind/website/page/group"
	profile_page "likemind/website/page/profile"
	signin_page "likemind/website/page/signin"
	signup_page "likemind/website/page/signup"
	user_search_page "likemind/website/page/user_search"
	"net/http"
)

type Server struct {
	auth    auth.Service
	profile profile.Service
}

var _ (desc.Handler) = (*Server)(nil)

func NewServer(auth auth.Service, profile profile.Service) *Server {
	return &Server{
		auth:    auth,
		profile: profile,
	}
}

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

func (s *Server) NotFound(w http.ResponseWriter, r *http.Request) {}
