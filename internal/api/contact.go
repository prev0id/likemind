package api

import (
	"context"

	"likemind/internal/common"
	"likemind/internal/domain"
	desc "likemind/internal/pkg/api"
)

func (s *Server) V1APIContactContactIDDelete(ctx context.Context, params desc.V1APIContactContactIDDeleteParams) (desc.V1APIContactContactIDDeleteRes, error) {
	userID := common.UserIDFromContext(ctx)

	err := s.profile.RemoveContact(ctx, userID, params.ContactID)
	if common.ErrorIs(err, common.NotFoundErrorType) {
		return &desc.NotFound{Data: common.ErrorMsg(err)}, nil
	}
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	// TODO:
	return &desc.HTMLResponse{}, nil
}

func (s *Server) V1APIContactContactIDPut(ctx context.Context, req *desc.Contact, params desc.V1APIContactContactIDPutParams) (desc.V1APIContactContactIDPutRes, error) {
	userID := common.UserIDFromContext(ctx)

	contact := convertContactToDomain(req)
	contact.ID = params.ContactID

	err := s.profile.UpdateContact(ctx, userID, contact)
	if common.ErrorIs(err, common.NotFoundErrorType) {
		return &desc.NotFound{Data: common.ErrorMsg(err)}, nil
	}
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	// TODO:
	return &desc.HTMLResponse{}, nil
}

func (s *Server) V1APIContactPost(ctx context.Context, req *desc.Contact) (desc.V1APIContactPostRes, error) {
	userID := common.UserIDFromContext(ctx)

	err := s.profile.AddContact(ctx, userID, convertContactToDomain(req))
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	// TODO:
	return &desc.HTMLResponse{}, nil
}

func convertContactToDomain(req *desc.Contact) domain.Contact {
	return domain.Contact{
		Platform: req.GetPlatform(),
		Value:    req.GetLink(),
	}
}
