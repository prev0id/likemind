package profile

import (
	"context"
	"fmt"

	"likemind/internal/database"
	"likemind/internal/database/adapter/profile_adapter"
	"likemind/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

const (
	saltPattern = "%d$%s"
	cost        = bcrypt.DefaultCost
)

type Service interface {
	CreateUser(ctx context.Context, user domain.User) (domain.UserID, error)
	UpdateUser(ctx context.Context, user domain.User) error
	DeleteProfile(ctx context.Context, id domain.UserID) error
	GetProfile(ctx context.Context, id domain.UserID) (domain.Profile, error)
	UpdatePassword(ctx context.Context, id domain.UserID, oldPassword, newPassword domain.Password) error
	SignIn(ctx context.Context, login domain.Email, password domain.Password) (domain.User, error)
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

	txErr := database.InTransaction(ctx, func(ctx context.Context) error {
		id, err = s.db.CreateUser(ctx, user)
		if err != nil {
			return fmt.Errorf("s.db.CreateUser: %w", err)
		}

		user.HashedPassword = hash(user.RawPassword, id)
		user.ID = id

		if err := s.db.UpdateUser(ctx, user); err != nil {
			return fmt.Errorf("s.db.UpdateUser: %w", err)
		}

		return nil
	})

	if txErr != nil {
		return 0, fmt.Errorf("database.InTransaction: %w", txErr)
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

func (s *implementation) GetProfile(ctx context.Context, id domain.UserID) (domain.Profile, error) {
	profile, err := s.db.GetProfileByUserID(ctx, id)
	if err != nil {
		return domain.Profile{}, fmt.Errorf("s.db.GetProfileByUserID: %w", err)
	}
	return profile, nil
}

func (s *implementation) UpdatePassword(ctx context.Context, id domain.UserID, oldPassword, newPassword domain.Password) error {
	user, err := s.db.GetUserByID(ctx, id)
	if err != nil {
		return fmt.Errorf("s.db.GetUserByID: %w", err)
	}

	if !passwordsEqual(user.HashedPassword, oldPassword, user.ID) {
		return domain.ErrNotAuthenticated
	}

	user.HashedPassword = hash(newPassword, id)

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

	if !passwordsEqual(user.HashedPassword, password, user.ID) {
		return domain.User{}, domain.ErrNotAuthenticated
	}

	return user, nil
}

func hash(password domain.Password, id domain.UserID) []byte {
	withSalt := addSalt(password, id)
	encrypted, _ := bcrypt.GenerateFromPassword(withSalt, cost)
	return encrypted
}

func passwordsEqual(hash []byte, password domain.Password, id domain.UserID) bool {
	withSalt := addSalt(password, id)
	err := bcrypt.CompareHashAndPassword(hash, withSalt)
	return err == nil
}

func addSalt(password domain.Password, id domain.UserID) []byte {
	return []byte(fmt.Sprintf(saltPattern, id, password))
}
