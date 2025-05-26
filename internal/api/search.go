package api

import (
	"context"
	"fmt"

	"likemind/internal/common"
	"likemind/website/view"
	"likemind/website/widget"

	desc "likemind/internal/pkg/api"

	"github.com/a-h/templ"
	"github.com/rs/zerolog/log"
)

func (s *Server) V1APISearchGet(ctx context.Context, params desc.V1APISearchGetParams) (desc.V1APISearchGetRes, error) {
	var result templ.Component

	userID := common.UserIDFromContext(ctx)

	switch params.Type {
	case desc.TypeProfile:
		userIDs, err := s.interests.SearchUsers(ctx, userID, params.Include, params.Exclude)
		if err != nil {
			return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
		}
		log.Info().Interface("ids", userID).Msg("api search")
		users := make([]*view.Profile, 0, len(userIDs))
		for _, userId := range userIDs {
			user, err := s.getProfile(ctx, userId)
			if err != nil {
				log.Err(err).Str("id", userID.String()).Msg("[api/search/user] s.getProfile")
				continue
			}
			users = append(users, user)
		}
		fmt.Println(users)
		result = widget.UserRecomendations(users)

	case desc.TypeGroup:
		groupIDs, err := s.interests.SearchGroups(ctx, userID, params.Include, params.Exclude)
		if err != nil {
			return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
		}
		groups := make([]*view.Group, 0, len(groupIDs))
		for _, groupID := range groupIDs {
			group, err := s.getGroup(ctx, groupID)
			if err != nil {
				log.Err(err).Str("id", groupID.String()).Msg("[api/search/group] s.getGroup")
				continue
			}
			groups = append(groups, group)
		}
		result = widget.GroupRecomendations(int64(userID), groups)
	}

	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, result),
	}, nil
}
