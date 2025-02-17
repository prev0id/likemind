package domain

import (
	"likemind/internal/common/validate"
)

type Profile struct {
	User            User
	ProfilePictures []string
	Contacts        []Contact
}

type Contact struct {
	ID       int64
	Platform string
	Value    string
}

type User struct {
	ID       int64
	Nickname string
	Name     Name
	Surname  Surname
	About    string
}

type Name string

func (n Name) Validate() error {
	return validate.String("name").
		NotEmpty().
		IsUTF8().
		LenMax(50).
		Build(string(n))
}

type Surname string

func (s Surname) Validate() error {
	return validate.String("surname").
		NotEmpty().
		IsUTF8().
		LenMax(50).
		Build(string(s))
}
