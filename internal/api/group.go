package api

import (
	"context"
	"fmt"
	"net/url"

	"likemind/internal/common"
	"likemind/internal/domain"
	"likemind/website/page"
	"likemind/website/widget"

	desc "likemind/internal/pkg/api"
)

func (s *Server) V1APIGroupPost(ctx context.Context, req *desc.Group) (desc.V1APIGroupPostRes, error) {
	fmt.Println("aaaaaaaaaaaaaaaaaaaa")

	group := domain.Group{
		Name:        req.GetName(),
		Description: req.GetDescription(),
	}

	id, err := s.group.CreateGroup(ctx, group)
	if common.ErrorIs(err, common.BadRequestErrorType) {
		return &desc.BadRequest{Data: common.ErrorMsg(err)}, nil
	}
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	if err := s.group.Subscribe(ctx, common.UserIDFromContext(ctx), id); err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	path := common.FillPath(
		domain.PathAPIGroupID,
		map[string]string{
			domain.PathParamContactID: id.String(),
		},
	)

	return &desc.Redirect302{
		HxRedirect: desc.NewOptURI(url.URL{Path: path}),
	}, nil
}

func (s *Server) V1APIGroupGroupIDDelete(
	ctx context.Context,
	params desc.V1APIGroupGroupIDDeleteParams,
) (desc.V1APIGroupGroupIDDeleteRes, error) {
	err := s.group.DeleteGroup(ctx, domain.GroupID(params.GroupID))
	if common.ErrorIs(err, common.PermissionDeniedErrorType) {
		return &desc.BadRequest{Data: common.ErrorMsg(err)}, nil
	}
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	return &desc.Redirect302{
		HxRedirect: desc.NewOptURI(url.URL{Path: domain.PathPageGroup}),
	}, nil
}

func (s *Server) V1APIGroupGroupIDPut(
	ctx context.Context,
	req *desc.Group,
	params desc.V1APIGroupGroupIDPutParams,
) (desc.V1APIGroupGroupIDPutRes, error) {
	group := domain.Group{
		ID:          domain.GroupID(params.GroupID),
		Name:        req.GetName(),
		Description: req.GetDescription(),
	}

	err := s.group.UpdateGroup(ctx, group)
	if common.ErrorIs(err, common.PermissionDeniedErrorType) {
		return &desc.NotAuthorized{Data: common.ErrorMsg(err)}, nil
	}
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	return &desc.HTMLResponse{
		// TODO:
	}, nil
}

func (s *Server) V1APIGroupGroupIDPostPost(
	ctx context.Context,
	req *desc.Post,
	params desc.V1APIGroupGroupIDPostPostParams,
) (desc.V1APIGroupGroupIDPostPostRes, error) {
	groupID := domain.GroupID(params.GroupID)
	post := domain.Post{
		Group:   groupID,
		Content: req.GetContent(),
	}

	_, err := s.group.CreatePost(ctx, post)
	if common.ErrorIs(err, common.BadRequestErrorType) {
		return &desc.BadRequest{Data: common.ErrorMsg(err)}, nil
	}
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	group, err := s.getGroup(ctx, groupID)
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, page.Posts(group)),
	}, nil
}

func (s *Server) V1APIGroupGroupIDPostPostIDCommentCommentIDDelete(
	ctx context.Context,
	params desc.V1APIGroupGroupIDPostPostIDCommentCommentIDDeleteParams,
) (desc.V1APIGroupGroupIDPostPostIDCommentCommentIDDeleteRes, error) {
	// TODO:
	return nil, nil
}

func (s *Server) V1APIGroupGroupIDPostPostIDCommentCommentIDPut(
	ctx context.Context,
	req *desc.Post,
	params desc.V1APIGroupGroupIDPostPostIDCommentCommentIDPutParams,
) (desc.V1APIGroupGroupIDPostPostIDCommentCommentIDPutRes, error) {
	// TODO:
	return nil, nil
}

func (s *Server) V1APIGroupGroupIDPostPostIDCommentPost(
	ctx context.Context,
	req *desc.Post,
	params desc.V1APIGroupGroupIDPostPostIDCommentPostParams,
) (desc.V1APIGroupGroupIDPostPostIDCommentPostRes, error) {
	groupID := domain.GroupID(params.GroupID)
	post := domain.Comment{
		PostID:  domain.PostID(params.PostID),
		Content: req.GetContent(),
	}

	_, err := s.group.CreateComment(ctx, post)
	if common.ErrorIs(err, common.BadRequestErrorType) {
		return &desc.BadRequest{Data: common.ErrorMsg(err)}, nil
	}
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	group, err := s.getGroup(ctx, groupID)
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, page.Posts(group)),
	}, nil
}

func (s *Server) V1APIGroupGroupIDPostPostIDDelete(
	ctx context.Context,
	params desc.V1APIGroupGroupIDPostPostIDDeleteParams,
) (desc.V1APIGroupGroupIDPostPostIDDeleteRes, error) {
	// TODO:
	return nil, nil
}

func (s *Server) V1APIGroupGroupIDPostPostIDPut(
	ctx context.Context,
	req *desc.Post,
	params desc.V1APIGroupGroupIDPostPostIDPutParams,
) (desc.V1APIGroupGroupIDPostPostIDPutRes, error) {
	// TODO
	return nil, nil
}

func (s *Server) V1APIGroupGroupIDSubscribeDelete(ctx context.Context, params desc.V1APIGroupGroupIDSubscribeDeleteParams) (desc.V1APIGroupGroupIDSubscribeDeleteRes, error) {
	userID := common.UserIDFromContext(ctx)
	groupID := domain.GroupID(params.GroupID)

	if err := s.group.Unsubscribe(ctx, userID, groupID); err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, widget.SubscribeButton(int64(groupID), false)),
	}, nil
}

func (s *Server) V1APIGroupGroupIDSubscribePost(ctx context.Context, params desc.V1APIGroupGroupIDSubscribePostParams) (desc.V1APIGroupGroupIDSubscribePostRes, error) {
	userID := common.UserIDFromContext(ctx)
	groupID := domain.GroupID(params.GroupID)

	if err := s.group.Subscribe(ctx, userID, groupID); err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, widget.SubscribeButton(int64(groupID), true)),
	}, nil
}
