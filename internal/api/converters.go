package api

import (
	"context"

	"likemind/internal/common"
	"likemind/internal/domain"
	"likemind/website/view"
)

func profileFromDomainToView(ctx context.Context, user domain.User, contacts []domain.Contact, pictures []domain.PictureID) *view.Profile {
	userID := common.UserIDFromContext(ctx)
	return &view.Profile{
		Name:        user.Name,
		Surname:     user.Surname,
		Nickname:    user.Nickname,
		About:       user.About,
		Location:    user.Location,
		DateOfBirth: user.DateOfBirth,
		Owner:       userID == user.ID,

		PictureID: convertPicturesIDs(pictures),
		Contacts:  contactsDomainToView(contacts),
	}
}

func convertPicturesIDs(picutres []domain.PictureID) string {
	if len(picutres) > 0 {
		return string(picutres[0])
	}
	return ""
}

func contactsDomainToView(contacts []domain.Contact) []view.Contact {
	result := make([]view.Contact, 0, len(contacts))
	for _, contact := range contacts {
		result = append(result, view.Contact{
			ID:       contact.ID,
			Platform: contact.Platform,
			Value:    contact.Value,
		})
	}
	return result
}
