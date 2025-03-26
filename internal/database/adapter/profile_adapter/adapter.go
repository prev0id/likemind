package profile_adapter

import (
	"context"
	"fmt"

	"likemind/internal/database"
	"likemind/internal/database/repo/contact_repo"
	profile_picture_repo "likemind/internal/database/repo/picture_repo"
	"likemind/internal/database/repo/user_repo"
	"likemind/internal/domain"
)

type Adapter interface {
	GetProfileByUserID(ctx context.Context, id domain.UserID) (domain.Profile, error)
	ListProfiles(ctx context.Context) ([]domain.Profile, error)

	CreateUser(ctx context.Context, user domain.User) (domain.UserID, error)
	UpdateUser(ctx context.Context, user domain.User) error
	RemoveUser(ctx context.Context, id domain.UserID) error
	GetUserByLogin(ctx context.Context, login domain.Email) (domain.User, error)
	GetUserByID(ctx context.Context, id domain.UserID) (domain.User, error)

	AddContact(ctx context.Context, id domain.UserID, contact domain.Contact) error
	UpdateContact(ctx context.Context, id domain.UserID, contact domain.Contact) error
	RemoveContactByID(ctx context.Context, contactID int64) error

	AddProfilePicture(ctx context.Context, id domain.UserID, pictureID string) error
	GetProfilePicturesByUserID(ctx context.Context, id domain.UserID) ([]string, error)
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

func (i *implementation) GetProfileByUserID(ctx context.Context, id domain.UserID) (domain.Profile, error) {
	user, err := i.userDB.GetUserByID(ctx, int64(id))
	if err != nil {
		return domain.Profile{}, fmt.Errorf("i.userDB.GetUserByID: %w", err)
	}

	contacts, err := i.contactDB.GetContactsByUserID(ctx, int64(id))
	if err != nil {
		return domain.Profile{}, fmt.Errorf("i.contactDB.GetContactsByUserID: %w", err)
	}

	pictures, err := i.pictureDB.GetProfilePicturesByUserID(ctx, int64(id))
	if err != nil {
		return domain.Profile{}, fmt.Errorf("i.pictureDB.GetProfilePicturesByUserID: %w", err)
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
		return nil, fmt.Errorf("i.userDB.ListUsers: %w", err)
	}

	profiles := make([]domain.Profile, 0, len(users))
	for _, user := range users {
		profile, err := i.GetProfileByUserID(ctx, domain.UserID(user.ID))
		if err != nil {
			return nil, fmt.Errorf("i.GetProfileByUserID: %w", err)
		}
		profiles = append(profiles, profile)
	}
	return profiles, nil
}

func (i *implementation) CreateUser(ctx context.Context, user domain.User) (domain.UserID, error) {
	id, err := i.userDB.CreateUser(ctx, domainUserToModel(user))
	if err != nil {
		return 0, fmt.Errorf("i.userDB.CreateUser: %w", err)
	}
	return domain.UserID(id), nil
}

func (i *implementation) UpdateUser(ctx context.Context, user domain.User) error {
	if err := i.userDB.UpdateUser(ctx, domainUserToModel(user)); err != nil {
		return fmt.Errorf("i.userDB.UpdateUser: %w", err)
	}
	return nil
}

func (i *implementation) RemoveUser(ctx context.Context, id domain.UserID) error {
	err := database.InTransaction(ctx, func(ctx context.Context) error {
		pictures, err := i.pictureDB.GetProfilePicturesByUserID(ctx, int64(id))
		if err != nil {
			return fmt.Errorf("i.pictureDB.GetProfilePicturesByUserID: %w", err)
		}
		for _, picture := range pictures {
			if err := i.pictureDB.RemovePictureByID(ctx, picture.ID); err != nil {
				return fmt.Errorf("i.pictureDB.RemovePictureByID: %w", err)
			}
		}

		contacts, err := i.contactDB.GetContactsByUserID(ctx, int64(id))
		if err != nil {
			return fmt.Errorf("i.contactDB.GetContactsByUserID: %w", err)
		}
		for _, contact := range contacts {
			if err := i.contactDB.RemoveContactByID(ctx, contact.ID); err != nil {
				return fmt.Errorf("i.contactDB.RemoveContactByID: %w", err)
			}
		}

		if err := i.userDB.RemoveUser(ctx, int64(id)); err != nil {
			return fmt.Errorf("i.userDB.RemoveUser: %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("RemoveUser transaction: %w", err)
	}
	return nil
}

func (i *implementation) GetUserByLogin(ctx context.Context, login domain.Email) (domain.User, error) {
	user, err := i.userDB.GetUserByEmail(ctx, string(login))
	if err != nil {
		return domain.User{}, fmt.Errorf("i.userDB.GetUserByLogin: %w", err)
	}
	return modelUserToDomain(user), nil
}

func (i *implementation) GetUserByID(ctx context.Context, id domain.UserID) (domain.User, error) {
	user, err := i.userDB.GetUserByID(ctx, int64(id))
	if err != nil {
		return domain.User{}, fmt.Errorf("i.userDB.GetUserByID: %w", err)
	}
	return modelUserToDomain(user), nil
}

func (i *implementation) AddContact(ctx context.Context, id domain.UserID, contact domain.Contact) error {
	if err := i.contactDB.AddContact(ctx, domainContactToModel(contact, int64(id))); err != nil {
		return fmt.Errorf("i.contactDB.AddContact: %w", err)
	}
	return nil
}

func (i *implementation) UpdateContact(ctx context.Context, id domain.UserID, contact domain.Contact) error {
	if err := i.contactDB.UpdateContact(ctx, domainContactToModel(contact, int64(id))); err != nil {
		return fmt.Errorf("i.contactDB.UpdateContact: %w", err)
	}
	return nil
}

func (i *implementation) RemoveContactByID(ctx context.Context, contactID int64) error {
	if err := i.contactDB.RemoveContactByID(ctx, contactID); err != nil {
		return fmt.Errorf("i.contactDB.RemoveContactByID: %w", err)
	}
	return nil
}

func (i *implementation) AddProfilePicture(ctx context.Context, id domain.UserID, pictureID string) error {
	if err := i.pictureDB.AddProfilePicture(ctx, domainProfilePictureToModel(pictureID, int64(id))); err != nil {
		return fmt.Errorf("i.pictureDB.AddProfilePicture: %w", err)
	}
	return nil
}

func (i *implementation) GetProfilePicturesByUserID(ctx context.Context, id domain.UserID) ([]string, error) {
	pictures, err := i.pictureDB.GetProfilePicturesByUserID(ctx, int64(id))
	if err != nil {
		return nil, fmt.Errorf("i.pictureDB.GetProfilePicturesByUserID: %w", err)
	}
	return convert(pictures, modelProfilePictureToDomain), nil
}

func (i *implementation) RemovePictureByID(ctx context.Context, pictureID string) error {
	if err := i.pictureDB.RemovePictureByID(ctx, pictureID); err != nil {
		return fmt.Errorf("i.pictureDB.RemovePictureByID: %w", err)
	}
	return nil
}
