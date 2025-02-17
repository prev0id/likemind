package profile_adapter

import (
	"likemind/internal/database/model"
	"likemind/internal/domain"
)

func convert[T, V any](data []T, converter func(T) V) []V {
	result := make([]V, 0, len(data))
	for _, element := range data {
		result = append(result, converter(element))
	}
	return result
}

func domainUserToModel(u domain.User) model.User {
	return model.User{
		ID:       u.ID,
		Nickname: u.Nickname,
		Name:     string(u.Name),
		Surname:  string(u.Surname),
		About:    u.About,
	}
}

func modelUserToDomain(u model.User) domain.User {
	return domain.User{
		ID:       u.ID,
		Nickname: u.Nickname,
		Name:     domain.Name(u.Name),
		Surname:  domain.Surname(u.Surname),
		About:    u.About,
	}
}

func domainContactToModel(c domain.Contact, userID int64) model.Contact {
	return model.Contact{
		ID:       c.ID,
		Platform: c.Platform,
		Value:    c.Value,
		UserID:   userID,
	}
}

func modelContactToDomain(c model.Contact) domain.Contact {
	return domain.Contact{
		ID:       c.ID,
		Platform: c.Platform,
		Value:    c.Value,
	}
}

func modelProfilePictureToDomain(pp model.ProfilePicture) string {
	return pp.ID
}

func domainProfilePictureToModel(pictureID string, userID int64) model.ProfilePicture {
	return model.ProfilePicture{
		ID:     pictureID,
		UserID: userID,
	}
}
