package domain

import (
	"fmt"

	"likemind/internal/common/validate"

	"golang.org/x/crypto/bcrypt"
)

const (
	saltPattern = "%d$%s"
	cost        = bcrypt.DefaultCost
)

type Credentials struct {
	ID       string
	UserID   int64
	Password []byte
	Login    string
}

type Password string

type Email string

func (p Password) Hash(userID int64) []byte {
	withSalt := p.addSalt(userID)
	encrypted, _ := bcrypt.GenerateFromPassword(withSalt, cost)
	return encrypted
}

func (p Password) CompareWithHash(hash []byte, userID int64) bool {
	withSalt := p.addSalt(userID)
	err := bcrypt.CompareHashAndPassword(hash, withSalt)
	return err == nil
}

func (p Password) addSalt(userID int64) []byte {
	withSalt := fmt.Sprintf(saltPattern, userID, p)
	return []byte(withSalt)
}

func (p Password) Validate() error {
	return validate.String("password").
		LenMin(6).
		LenMax(20).
		ContainsSymbol().
		ContainsDigit().
		ContainsUpper().
		ContainsLower().
		Build(string(p))
}

func (e Email) Validate() error {
	return validate.String("login").
		NotEmpty().
		LenMax(50).
		Pattern(validate.PatternEmail).
		Build(string(e))
}
