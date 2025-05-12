package api

import (
	"likemind/internal/domain"
	"likemind/website/view"
)

func profileFromDomainToView(
	userID domain.UserID,
	user domain.User,
	contacts []domain.Contact,
	pictures []domain.PictureID,
	interests []domain.InterestGroup,
) *view.Profile {
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

		Interests: interestGroupDomainToView(interests),
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

func interestGroupDomainToView(interests []domain.InterestGroup) []view.GroupedInterests {
	result := make([]view.GroupedInterests, 0, len(interests))
	for _, group := range interests {
		result = append(result, view.GroupedInterests{
			Name:      group.Name,
			Interests: interestsDomainToView(group.Interests),
		})
	}
	return result
}

func interestsDomainToView(interests []domain.Interest) []view.Interest {
	result := make([]view.Interest, 0, len(interests))
	for _, interest := range interests {
		result = append(result, view.Interest{
			ID:          int64(interest.ID),
			Name:        interest.Name,
			Description: interest.Description,
			Selected:    interest.Selected,
		})
	}
	return result
}
