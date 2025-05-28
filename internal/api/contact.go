package api

import (
	"context"
	"fmt"

	"likemind/internal/common"
	"likemind/internal/domain"
	"likemind/website/widget"

	desc "likemind/internal/pkg/api"
)

func (s *Server) V1APIProfileContactContactIDDelete(ctx context.Context, params desc.V1APIProfileContactContactIDDeleteParams) (desc.V1APIProfileContactContactIDDeleteRes, error) {
	userID := common.UserIDFromContext(ctx)
	fmt.Println(userID, params.ContactID)

	err := s.profile.RemoveContact(ctx, userID, params.ContactID)
	if common.ErrorIs(err, common.NotFoundErrorType) {
		return &desc.NotFound{Data: common.ErrorMsg(err)}, nil
	}
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	contacts, err := s.profile.GetContacts(ctx, userID)
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	component := widget.UpdateContacts(contactsDomainToView(contacts))

	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, component),
	}, nil
}

func (s *Server) V1APIProfileContactContactIDPut(ctx context.Context, req *desc.Contact, params desc.V1APIProfileContactContactIDPutParams) (desc.V1APIProfileContactContactIDPutRes, error) {
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

	contacts, err := s.profile.GetContacts(ctx, userID)
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	component := widget.UpdateContacts(contactsDomainToView(contacts))

	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, component),
	}, nil
}

func (s *Server) V1APIProfileContactPost(ctx context.Context, req *desc.Contact) (desc.V1APIProfileContactPostRes, error) {
	userID := common.UserIDFromContext(ctx)

	err := s.profile.AddContact(ctx, userID, convertContactToDomain(req))
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	contacts, err := s.profile.GetContacts(ctx, userID)
	if err != nil {
		return &desc.InternalError{Data: common.ErrorMsg(err)}, nil
	}

	component := widget.UpdateContacts(contactsDomainToView(contacts))

	return &desc.HTMLResponse{
		Data: common.RenderComponent(ctx, component),
	}, nil
}

func convertContactToDomain(req *desc.Contact) domain.Contact {
	return domain.Contact{
		Platform: req.GetPlatform(),
		Value:    req.GetLink(),
	}
}
