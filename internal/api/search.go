package api

import (
	"context"

	"likemind/internal/common"
	"likemind/website/view"
	"likemind/website/widget"

	desc "likemind/internal/pkg/api"

	"github.com/a-h/templ"
)

func (s *Server) V1APISearchGet(ctx context.Context, req *desc.Search) (desc.V1APISearchGetRes, error) {
	var result templ.Component

	userID := common.UserIDFromContext(ctx)

	switch req.GetType() {
	case desc.SearchTypeProfile:
		userIDs, err := s.interests.SearchUsers(ctx, userID, req.IncludeInterests, req.ExcludeInterests)
		if err != nil {
			return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
		}
		users := make([]*view.Profile, 0, len(userIDs))
		for _, userId := range userIDs {
			user, err := s.getProfile(ctx, userId)
			if err != nil {
				continue
			}
			users = append(users, user)
		}
		result = widget.UserRecomendations(users)

	case desc.SearchTypeGroup:
		groupIDs, err := s.interests.SearchGroups(ctx, userID, req.IncludeInterests, req.ExcludeInterests)
		if err != nil {
			return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
		}
		groups := make([]*view.Group, 0, len(groupIDs))
		for _, groupID := range groupIDs {
			group, err := s.getGroup(ctx, groupID)
			if err != nil {
				continue
			}
			groups = append(groups, group)
		}
		result = widget.GroupRecomendations(groups)
	}

	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, result),
	}, nil
}
