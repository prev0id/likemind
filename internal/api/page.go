package api

import (
	"context"
	"likemind/internal/common"
	"likemind/internal/domain"
	"likemind/website/page"
	"likemind/website/view"
	"likemind/website/widget"
	"net/http"

	desc "likemind/internal/pkg/api"

	error_page "likemind/website/page/error"
)

func (s *Server) V1PageGroupGet(ctx context.Context) (desc.V1PageGroupGetRes, error) {
	userID := common.UserIDFromContext(ctx)

	subscriptions, err := s.group.ListSubscribedGroups(ctx, userID)
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	groups := make([]*view.Group, 0, len(subscriptions))
	for _, sub := range subscriptions {
		group, err := s.getGroup(ctx, sub)
		if err != nil {
			continue
		}
		groups = append(groups, group)
	}

	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, page.GroupSubscriptions(groups)),
	}, nil
}

func (s *Server) V1PageGroupGroupIDGet(ctx context.Context, params desc.V1PageGroupGroupIDGetParams) (desc.V1PageGroupGroupIDGetRes, error) {
	group, err := s.getGroup(ctx, domain.GroupID(params.GroupID))
	if err != nil {
		if common.ErrorIs(err, common.NotFoundErrorType) {
			return &desc.NotFound{Data: common.ErrorMsg(err)}, nil
		}
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, page.Group(group)),
	}, nil
}

func (s *Server) V1PageProfileUsernameGet(ctx context.Context, params desc.V1PageProfileUsernameGetParams) (desc.V1PageProfileUsernameGetRes, error) {
	userID := common.UserIDFromContext(ctx)

	profile, err := s.getProfile(ctx, userID)
	if err != nil {
		if common.ErrorIs(err, common.NotFoundErrorType) {
			return &desc.NotFound{Data: common.ErrorMsg(err)}, nil
		}
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	if int64(userID) == profile.ID {
		return &desc.Redirect302{
			Location: getProfilePage(),
		}, nil
	}

	pageComponent := page.Profile(profile)

	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, pageComponent),
	}, nil
}

func (s *Server) V1PageProfileGet(ctx context.Context) (desc.V1PageProfileGetRes, error) {
	userID := common.UserIDFromContext(ctx)

	profile, err := s.getProfile(ctx, userID)
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	pageComponent := page.Profile(profile)

	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, pageComponent),
	}, nil
}

func (s *Server) V1PageSearchGet(ctx context.Context) (desc.V1PageSearchGetRes, error) {
	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, page.Search(widget.Placeholder())),
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
