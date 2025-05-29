package api

import (
	"context"
	"net/http"

	"likemind/internal/common"
	"likemind/internal/domain"
	desc "likemind/internal/pkg/api"
	"likemind/website/page"
	error_page "likemind/website/page/error"
	"likemind/website/view"
	"likemind/website/widget"

	"github.com/rs/zerolog/log"
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
		Data: common.RenderComponent(ctx, page.GroupSubscriptions(int64(userID), groups)),
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
		Data: common.RenderComponent(ctx, page.Group(group, int64(common.UserIDFromContext(ctx)))),
	}, nil
}

func (s *Server) V1PageProfileUsernameGet(ctx context.Context, params desc.V1PageProfileUsernameGetParams) (desc.V1PageProfileUsernameGetRes, error) {
	userID := common.UserIDFromContext(ctx)

	profile, err := s.getProfileByUsername(ctx, params.Username)
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
	userID := common.UserIDFromContext(ctx)

	userIDs, err := s.interests.SearchUsers(ctx, userID, nil, nil)
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	log.Info().Interface("ids", userIDs).Int("size", len(userIDs)).Msg("page search")

	users := make([]*view.Profile, 0, len(userIDs))
	for _, userId := range userIDs {
		user, err := s.getProfile(ctx, userId)
		if err != nil {
			log.Err(err).Str("id", userId.String()).Msg("[page/search] s.getProfile")
			continue
		}
		users = append(users, user)
	}

	recomendations := widget.UserRecomendations(users)

	interests, err := s.interests.ListInterests(ctx)
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, page.Search(interestGroupDomainToView(interests), recomendations)),
	}, nil
}

func (s *Server) V1PageSigninGet(ctx context.Context, params desc.V1PageSigninGetParams) (desc.V1PageSigninGetRes, error) {
	if common.UserIDFromContext(ctx) > 0 {
		return &desc.Redirect302{
			Location: getProfilePage(),
		}, nil
	}

	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, page.SignIn()),
	}, nil
}

func (s *Server) V1PageSignupGet(ctx context.Context, params desc.V1PageSignupGetParams) (desc.V1PageSignupGetRes, error) {
	if common.UserIDFromContext(ctx) > 0 {
		return &desc.Redirect302{
			Location: getProfilePage(),
		}, nil
	}

	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, page.SignUp()),
	}, nil
}

func (s *Server) NotFound(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/v1/page/" || r.URL.Path == "/v1/page" {
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}

	page := error_page.Page(error_page.State{Code: http.StatusNotFound})

	w.WriteHeader(http.StatusNotFound)
	page.Render(r.Context(), w)
}
