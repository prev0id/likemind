package auth

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	"likemind/internal/domain"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"golang.org/x/crypto/bcrypt"
)

func (i *implementation) SetCredentials(ctx context.Context, userID int64, login, password string) (string, error) {
	creds := domain.Credential{
		UserID:   userID,
		Login:    login,
		Password: encryptPassword(userID, password),
		UUID:     uuid.NewString(),
	}

	_, err := i.db.Insert(ctx, creds)
	if err != nil {
		return "", fmt.Errorf("i.db.Insert: %w", err)
	}

	return creds.UUID, nil
}

func (i *implementation) ValidateCredentials(ctx context.Context, login string, password string) (string, error) {
	creds, err := i.findCredentialByLogin(ctx, login)
	if err != nil {
		return "", err
	}

	if !validatePassword(creds.Password, creds.UserID, password) {
		return "", errors.New("incorrect password")
	}

	return creds.UUID, nil
}

func (i *implementation) findCredentialByLogin(ctx context.Context, login string) (domain.Credential, error) {
	creds, err := i.db.List(ctx)
	if err != nil {
		return domain.Credential{}, fmt.Errorf("i.db.List: %w", err)
	}

	cred, found := lo.Find(creds, func(cred domain.Credential) bool {
		return cred.Login == login
	})

	if !found {
		return domain.Credential{}, fmt.Errorf("user with login '%s' not found", login)
	}

	return cred, nil
}

func validateNewPassword(login, password string) error {
	if len(password) < passwordMinLen {
		return fmt.Errorf("password is to short (min lenght is %d", passwordMinLen)
	}

	if len(password) > passwordMaxLen {
		return fmt.Errorf("password is to long (max lenght is %d", passwordMaxLen)
	}

	if !utf8.ValidString(password) {
		return errors.New("password should be valid UTF-8 string")
	}

	if !strings.ContainsAny(password, numbers) {
		return errors.New("password must contain any number")
	}

	if !strings.ContainsAny(password, specialChars) {
		return fmt.Errorf("password must contain any of specital symbols %s", specialChars)
	}

	if strings.Contains(password, login) {
		return errors.New("password must not contain your login")
	}

	if strings.Contains(login, password) {
		return errors.New("password must not be part of your login")
	}

	return nil
}

func encryptPassword(userID int64, password string) []byte {
	withSalt := passwordWithSalt(userID, password)
	encrypted, _ := bcrypt.GenerateFromPassword(withSalt, bcrypt.DefaultCost)
	return encrypted
}

func validatePassword(encrypted []byte, userID int64, password string) bool {
	withSalt := passwordWithSalt(userID, password)
	err := bcrypt.CompareHashAndPassword(encrypted, withSalt)
	return err == nil
}

func passwordWithSalt(userID int64, password string) []byte {
	result := fmt.Sprintf(saltPattern, userID, password)
	return []byte(result)
}
