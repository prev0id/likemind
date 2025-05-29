package profile

import (
	"context"
	"fmt"

	"likemind/internal/database/adapter/profile_adapter"
	"likemind/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

const (
	saltPattern = "%s$%s"
	cost        = bcrypt.DefaultCost
)

type Service interface {
	DeleteProfile(ctx context.Context, id domain.UserID) error

	CreateUser(ctx context.Context, user domain.User) (domain.UserID, error)
	UpdateUser(ctx context.Context, user domain.User) error
	GetUser(ctx context.Context, id domain.UserID) (domain.User, error)
	UpdatePassword(ctx context.Context, id domain.UserID, login domain.Email, oldPassword, newPassword domain.Password) error
	UpdateEmail(ctx context.Context, id domain.UserID, oldLogin, login domain.Email, password domain.Password) error
	SignIn(ctx context.Context, login domain.Email, password domain.Password) (domain.User, error)
	GetUserByUsername(ctx context.Context, username string) (domain.User, error)

	GetContacts(ctx context.Context, id domain.UserID) ([]domain.Contact, error)
	AddContact(ctx context.Context, id domain.UserID, contact domain.Contact) error
	UpdateContact(ctx context.Context, id domain.UserID, contact domain.Contact) error
	RemoveContact(ctx context.Context, id domain.UserID, contactID int64) error
}

type implementation struct {
	db profile_adapter.Adapter
}

func New(db profile_adapter.Adapter) Service {
	return &implementation{
		db: db,
	}
}

func (s *implementation) CreateUser(ctx context.Context, user domain.User) (domain.UserID, error) {
	var (
		id  domain.UserID
		err error
	)

	user.HashedPassword = hash(user.RawPassword, user.Login)

	id, err = s.db.CreateUser(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("s.db.CreateUser: %w", err)
	}

	return id, nil
}

func (s *implementation) UpdateUser(ctx context.Context, user domain.User) error {
	if err := s.db.UpdateUser(ctx, user); err != nil {
		return fmt.Errorf("s.db.UpdateUser: %w", err)
	}

	return nil
}

func (s *implementation) DeleteProfile(ctx context.Context, id domain.UserID) error {
	if err := s.db.RemoveUser(ctx, id); err != nil {
		return fmt.Errorf("s.db.RemoveUser: %w", err)
	}

	return nil
}

func (s *implementation) GetUser(ctx context.Context, id domain.UserID) (domain.User, error) {
	user, err := s.db.GetUserByID(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("s.db.GetUserByID: %w", err)
	}

	return user, nil
}

func (s *implementation) UpdatePassword(ctx context.Context, id domain.UserID, email domain.Email, oldPassword, newPassword domain.Password) error {
	user, err := s.db.GetUserByID(ctx, id)
	if err != nil {
		return fmt.Errorf("s.db.GetUserByID: %w", err)
	}

	if !passwordsEqual(user.HashedPassword, oldPassword, email) {
		return domain.ErrNotAuthenticated
	}

	user.HashedPassword = hash(newPassword, email)

	if err := s.db.UpdateUser(ctx, user); err != nil {
		return fmt.Errorf("s.db.UpdateUser: %w", err)
	}

	return nil
}

func (s *implementation) UpdateEmail(ctx context.Context, id domain.UserID, oldEmail, newEmail domain.Email, password domain.Password) error {
	user, err := s.db.GetUserByID(ctx, id)
	if err != nil {
		return fmt.Errorf("s.db.GetUserByID: %w", err)
	}

	if !passwordsEqual(user.HashedPassword, password, oldEmail) {
		return domain.ErrNotAuthenticated
	}

	user.Login = newEmail
	user.HashedPassword = hash(password, newEmail)

	if err := s.db.UpdateUser(ctx, user); err != nil {
		return fmt.Errorf("s.db.UpdateUser: %w", err)
	}

	return nil
}

func (s *implementation) SignIn(ctx context.Context, login domain.Email, password domain.Password) (domain.User, error) {
	user, err := s.db.GetUserByLogin(ctx, login)
	if err != nil {
		return domain.User{}, fmt.Errorf("s.db.GetUserByLogin: %w", err)
	}

	if !passwordsEqual(user.HashedPassword, password, login) {
		return domain.User{}, domain.ErrNotAuthenticated
	}

	return user, nil
}

func hash(password domain.Password, email domain.Email) []byte {
	withSalt := addSalt(password, email)
	encrypted, _ := bcrypt.GenerateFromPassword(withSalt, cost)
	return encrypted
}

func passwordsEqual(hash []byte, password domain.Password, email domain.Email) bool {
	withSalt := addSalt(password, email)
	err := bcrypt.CompareHashAndPassword(hash, withSalt)
	return err == nil
}

func addSalt(password domain.Password, email domain.Email) []byte {
	return fmt.Appendf(nil, saltPattern, email, password)
}

func (s *implementation) GetUserByUsername(ctx context.Context, username string) (domain.User, error) {
	user, err := s.db.GetUserByUsername(ctx, username)
	if err != nil {
		return domain.User{}, fmt.Errorf("s.db.GetUserByUsername: %w", err)
	}
	return user, nil
}
