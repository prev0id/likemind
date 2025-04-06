package profile_adapter

import (
	"context"
	"fmt"

	"likemind/internal/common"
	"likemind/internal/database"
	"likemind/internal/database/repo/contact_repo"
	profile_picture_repo "likemind/internal/database/repo/picture_repo"
	"likemind/internal/database/repo/user_repo"
	"likemind/internal/domain"
)

type Adapter interface {
	CreateUser(ctx context.Context, user domain.User) (domain.UserID, error)
	UpdateUser(ctx context.Context, user domain.User) error
	RemoveUser(ctx context.Context, id domain.UserID) error
	GetUserByLogin(ctx context.Context, login domain.Email) (domain.User, error)
	GetUserByID(ctx context.Context, id domain.UserID) (domain.User, error)

	AddContact(ctx context.Context, id domain.UserID, contact domain.Contact) error
	UpdateContact(ctx context.Context, id domain.UserID, contact domain.Contact) error
	RemoveContactByID(ctx context.Context, contactID int64) error
	GetContactsByUserID(ctx context.Context, id domain.UserID) ([]domain.Contact, error)

	AddProfilePicture(ctx context.Context, id domain.UserID, pictureID domain.PictureID) error
	GetProfilePicturesByUserID(ctx context.Context, id domain.UserID) ([]domain.PictureID, error)
	RemovePictureByID(ctx context.Context, pictureID domain.PictureID) error
}

type Implementation struct {
	user    user_repo.DB
	contact contact_repo.DB
	picture profile_picture_repo.DB
}

var _ Adapter = (*Implementation)(nil)

func NewAdapter(userDB user_repo.DB, contactDB contact_repo.DB, pictureDB profile_picture_repo.DB) *Implementation {
	return &Implementation{
		user:    userDB,
		contact: contactDB,
		picture: pictureDB,
	}
}

func (i *Implementation) CreateUser(ctx context.Context, user domain.User) (domain.UserID, error) {
	id, err := i.user.Create(ctx, domainUserToModel(user))
	if err != nil {
		return 0, fmt.Errorf("i.userDB.CreateUser: %w", err)
	}
	return domain.UserID(id), nil
}

func (i *Implementation) UpdateUser(ctx context.Context, user domain.User) error {
	if err := i.user.Update(ctx, domainUserToModel(user)); err != nil {
		return fmt.Errorf("i.userDB.UpdateUser: %w", err)
	}
	return nil
}

func (i *Implementation) RemoveUser(ctx context.Context, id domain.UserID) error {
	err := database.InTransaction(ctx, func(ctx context.Context) error {
		pictures, err := i.picture.GetProfilePicturesByUserID(ctx, int64(id))
		if err != nil {
			return fmt.Errorf("i.pictureDB.GetProfilePicturesByUserID: %w", err)
		}
		for _, picture := range pictures {
			if pictureErr := i.picture.RemovePictureByID(ctx, picture.ID); pictureErr != nil {
				return fmt.Errorf("i.pictureDB.RemovePictureByID: %w", pictureErr)
			}
		}

		contacts, err := i.contact.GetContactsByUserID(ctx, int64(id))
		if err != nil {
			return fmt.Errorf("i.contactDB.GetContactsByUserID: %w", err)
		}
		for _, contact := range contacts {
			if err := i.contact.RemoveContactByID(ctx, contact.ID); err != nil {
				return fmt.Errorf("i.contactDB.RemoveContactByID: %w", err)
			}
		}

		if err := i.user.Delete(ctx, int64(id)); err != nil {
			return fmt.Errorf("i.userDB.RemoveUser: %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("RemoveUser transaction: %w", err)
	}
	return nil
}

func (i *Implementation) GetUserByLogin(ctx context.Context, login domain.Email) (domain.User, error) {
	user, err := i.user.GetByEmail(ctx, string(login))
	if err != nil {
		return domain.User{}, fmt.Errorf("i.userDB.GetUserByLogin: %w", err)
	}
	return modelUserToDomain(user), nil
}

func (i *Implementation) GetUserByID(ctx context.Context, id domain.UserID) (domain.User, error) {
	user, err := i.user.GetByID(ctx, int64(id))
	if err != nil {
		return domain.User{}, fmt.Errorf("i.userDB.GetUserByID: %w", err)
	}
	return modelUserToDomain(user), nil
}

func (i *Implementation) AddContact(ctx context.Context, id domain.UserID, contact domain.Contact) error {
	if err := i.contact.AddContact(ctx, domainContactToModel(contact, int64(id))); err != nil {
		return fmt.Errorf("i.contactDB.AddContact: %w", err)
	}
	return nil
}

func (i *Implementation) UpdateContact(ctx context.Context, id domain.UserID, contact domain.Contact) error {
	if err := i.contact.UpdateContact(ctx, domainContactToModel(contact, int64(id))); err != nil {
		return fmt.Errorf("i.contactDB.UpdateContact: %w", err)
	}
	return nil
}

func (i *Implementation) RemoveContactByID(ctx context.Context, contactID int64) error {
	if err := i.contact.RemoveContactByID(ctx, contactID); err != nil {
		return fmt.Errorf("i.contactDB.RemoveContactByID: %w", err)
	}
	return nil
}

func (i *Implementation) AddProfilePicture(ctx context.Context, id domain.UserID, pictureID domain.PictureID) error {
	if err := i.picture.AddProfilePicture(ctx, domainProfilePictureToModel(pictureID, int64(id))); err != nil {
		return fmt.Errorf("i.pictureDB.AddProfilePicture: %w", err)
	}
	return nil
}

func (i *Implementation) GetProfilePicturesByUserID(ctx context.Context, id domain.UserID) ([]domain.PictureID, error) {
	pictures, err := i.picture.GetProfilePicturesByUserID(ctx, int64(id))
	if err != nil {
		return nil, fmt.Errorf("i.pictureDB.GetProfilePicturesByUserID: %w", err)
	}
	return common.Convert(pictures, modelProfilePictureToDomain), nil
}

func (i *Implementation) RemovePictureByID(ctx context.Context, pictureID domain.PictureID) error {
	if err := i.picture.RemovePictureByID(ctx, string(pictureID)); err != nil {
		return fmt.Errorf("i.pictureDB.RemovePictureByID: %w", err)
	}
	return nil
}

func (i *Implementation) GetContactsByUserID(ctx context.Context, id domain.UserID) ([]domain.Contact, error) {
	contacts, err := i.contact.GetContactsByUserID(ctx, int64(id))
	if err != nil {
		return nil, fmt.Errorf("i.contactDB.GetContactsByUserID: %w", err)
	}

	return common.Convert(contacts, modelContactToDomain), nil
}
