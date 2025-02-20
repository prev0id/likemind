package auth

import (
	"context"
	"fmt"

	"likemind/internal/domain"

	"github.com/google/uuid"
)

func (i *implementation) NewUserCredentials(ctx context.Context, userID int64, login domain.Email, password domain.Password) (string, error) {
	if err := password.Validate(); err != nil {
		return "", err
	}

	if err := login.Validate(); err != nil {
		return "", err
	}

	creds := domain.Credentials{
		ID:       uuid.NewString(),
		UserID:   userID,
		Login:    string(login),
		Password: password.Hash(userID),
	}

	if err := i.db.Create(ctx, creds); err != nil {
		return "", fmt.Errorf("i.db.Insert: %w", err)
	}

	return creds.ID, nil
}

func (i *implementation) Signin(ctx context.Context, login domain.Email, password domain.Password) (domain.Credentials, error) {
	creds, err := i.db.GetByLogin(ctx, login)
	if err != nil {
		return domain.Credentials{}, fmt.Errorf("i.db.Get: %w", err)
	}

	if !password.CompareWithHash(creds.Password, creds.UserID) {
		return domain.Credentials{}, domain.ErrUnauthenticated
	}

	return creds, nil
}

func (i *implementation) Authenticate(ctx context.Context, credsID string) (int64, error) {
	creds, err := i.db.GetByID(ctx, credsID)
	if err != nil {
		return 0, domain.ErrUnauthenticated
	}

	return creds.UserID, nil
}

func (i *implementation) UpdatePassword(ctx context.Context, id string, newPassword domain.Password) error {
	creds, err := i.db.GetByID(ctx, id)
	if err != nil {
		return err
	}

	creds.Password = newPassword.Hash(creds.UserID)

	if err := i.db.Update(ctx, creds); err != nil {
		return err
	}

	return nil
}
