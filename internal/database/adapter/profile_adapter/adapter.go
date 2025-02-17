package profile_adapter

import (
	"context"

	"likemind/internal/database/repo/contact_repo"
	profile_picture_repo "likemind/internal/database/repo/picture_repo"
	"likemind/internal/database/repo/user_repo"
	"likemind/internal/domain"
)

type Adapter interface {
	GetProfileByUserID(ctx context.Context, userID int64) (domain.Profile, error)
	ListProfiles(ctx context.Context) ([]domain.Profile, error)

	CreateUser(ctx context.Context, user domain.User) error
	UpdateUser(ctx context.Context, user domain.User) error

	AddContact(ctx context.Context, userID int64, contact domain.Contact) error
	UpdateContact(ctx context.Context, userID int64, contact domain.Contact) error
	RemoveContactByID(ctx context.Context, contactID int64) error

	AddProfilePicutre(ctx context.Context, userID int64, pictureID string) error
	GetProfilePicutresByUserID(ctx context.Context, userID int64) ([]string, error)
	RemovePictureByID(ctx context.Context, pictureID string) error
}

type implementation struct {
	userDB    user_repo.DB
	contactDB contact_repo.DB
	pictureDB profile_picture_repo.DB
}

func NewAdapter(userDB user_repo.DB, contactDB contact_repo.DB, pictureDB profile_picture_repo.DB) Adapter {
	return &implementation{
		userDB:    userDB,
		contactDB: contactDB,
		pictureDB: pictureDB,
	}
}

func (i *implementation) GetProfileByUserID(ctx context.Context, userID int64) (domain.Profile, error) {
	user, err := i.userDB.GetUserByID(ctx, userID)
	if err != nil {
		return domain.Profile{}, err
	}

	contacts, err := i.contactDB.GetContactsByUserID(ctx, userID)
	if err != nil {
		return domain.Profile{}, err
	}

	pictures, err := i.pictureDB.GetProfilePicutresByUserID(ctx, userID)
	if err != nil {
		return domain.Profile{}, err
	}

	return domain.Profile{
		User:            modelUserToDomain(user),
		Contacts:        convert(contacts, modelContactToDomain),
		ProfilePictures: convert(pictures, modelProfilePictureToDomain),
	}, nil
}

func (i *implementation) ListProfiles(ctx context.Context) ([]domain.Profile, error) {
	users, err := i.userDB.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	profiles := make([]domain.Profile, 0, len(users))
	for _, userModel := range users {
		profile, err := i.GetProfileByUserID(ctx, userModel.ID)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}
	return profiles, nil
}

func (i *implementation) CreateUser(ctx context.Context, user domain.User) error {
	return i.userDB.CreateUser(ctx, domainUserToModel(user))
}

func (i *implementation) UpdateUser(ctx context.Context, user domain.User) error {
	return i.userDB.UpdateUser(ctx, domainUserToModel(user))
}

func (i *implementation) AddContact(ctx context.Context, userID int64, contact domain.Contact) error {
	return i.contactDB.AddContact(ctx, domainContactToModel(contact, userID))
}

func (i *implementation) UpdateContact(ctx context.Context, userID int64, contact domain.Contact) error {
	return i.contactDB.UpdateContact(ctx, domainContactToModel(contact, userID))
}

func (i *implementation) RemoveContactByID(ctx context.Context, contactID int64) error {
	return i.contactDB.RemoveContactByID(ctx, contactID)
}

func (i *implementation) AddProfilePicutre(ctx context.Context, userID int64, pictureID string) error {
	return i.pictureDB.AddProfilePicutre(ctx, domainProfilePictureToModel(pictureID, userID))
}

func (i *implementation) GetProfilePicutresByUserID(ctx context.Context, userID int64) ([]string, error) {
	pictures, err := i.pictureDB.GetProfilePicutresByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return convert(pictures, modelProfilePictureToDomain), nil
}

func (i *implementation) RemovePictureByID(ctx context.Context, pictureID string) error {
	return i.pictureDB.RemovePictureByID(ctx, pictureID)
}
