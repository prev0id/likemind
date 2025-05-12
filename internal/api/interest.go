package api

import (
	"context"

	"likemind/internal/common"
	"likemind/internal/domain"
	desc "likemind/internal/pkg/api"
	"likemind/website/widget"
)

func (s *Server) V1APIProfileInterestInterestIDDelete(ctx context.Context, params desc.V1APIProfileInterestInterestIDDeleteParams) (desc.V1APIProfileInterestInterestIDDeleteRes, error) {
	userID := common.UserIDFromContext(ctx)

	interests, err := s.interests.DeleteInterestFromUser(ctx, userID, domain.InterestID(params.InterestID))
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	form := widget.SelectInterests(interestGroupDomainToView(interests), domain.PathAPIProfileInterestID)

	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, form),
	}, nil
}

func (s *Server) V1APIProfileInterestInterestIDPost(ctx context.Context, params desc.V1APIProfileInterestInterestIDPostParams) (desc.V1APIProfileInterestInterestIDPostRes, error) {
	userID := common.UserIDFromContext(ctx)

	interests, err := s.interests.AddInterestToUser(ctx, userID, domain.InterestID(params.InterestID))
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	form := widget.SelectInterests(interestGroupDomainToView(interests), domain.PathAPIProfileInterestID)

	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, form),
	}, nil
}

func (s *Server) V1APIGroupGroupIDInterestInterestIDDelete(ctx context.Context, params desc.V1APIGroupGroupIDInterestInterestIDDeleteParams) (desc.V1APIGroupGroupIDInterestInterestIDDeleteRes, error) {
	interests, err := s.interests.DeleteInterestFromGroup(ctx, domain.GroupID(params.GroupID), domain.InterestID(params.InterestID))
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	form := widget.SelectInterests(interestGroupDomainToView(interests), domain.PathAPIGroupInterestID)

	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, form),
	}, nil
}

func (s *Server) V1APIGroupGroupIDInterestInterestIDPost(ctx context.Context, params desc.V1APIGroupGroupIDInterestInterestIDPostParams) (desc.V1APIGroupGroupIDInterestInterestIDPostRes, error) {
	interests, err := s.interests.AddInterestToGroup(ctx, domain.GroupID(params.GroupID), domain.InterestID(params.InterestID))
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	form := widget.SelectInterests(interestGroupDomainToView(interests), domain.PathAPIGroupInterestID)

	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, form),
	}, nil
}
