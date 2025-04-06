package profile_adapter

import (
	"likemind/internal/database/model"
	"likemind/internal/domain"
)

func domainUserToModel(u domain.User) model.User {
	return model.User{
		ID:       int64(u.ID),
		Nickname: u.Nickname,
		Name:     string(u.Name),
		Surname:  string(u.Surname),
		About:    u.About,
		Password: u.HashedPassword,
		Email:    string(u.Login),
		Location: u.Location,
	}
}

func modelUserToDomain(u model.User) domain.User {
	return domain.User{
		ID:             domain.UserID(u.ID),
		Nickname:       u.Nickname,
		Name:           u.Name,
		Surname:        u.Surname,
		About:          u.About,
		HashedPassword: u.Password,
		Login:          domain.Email(u.About),
		Location:       u.Location,
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

func modelProfilePictureToDomain(pp model.ProfilePicture) domain.PictureID {
	return domain.PictureID(pp.ID)
}

func domainProfilePictureToModel(pictureID domain.PictureID, userID int64) model.ProfilePicture {
	return model.ProfilePicture{
		ID:     string(pictureID),
		UserID: userID,
	}
}
