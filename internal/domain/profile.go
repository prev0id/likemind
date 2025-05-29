package domain

import (
	"strconv"
)

type (
	UserID    int64
	Password  string
	Email     string
	PictureID string
)

type User struct {
	ID             UserID
	Nickname       string
	Name           string
	Surname        string
	About          string
	Location       string
	Login          Email
	HashedPassword []byte
	RawPassword    Password
}

type Contact struct {
	ID       int64
	Platform string
	Value    string
}

func (id UserID) String() string {
	return strconv.FormatInt(int64(id), 10)
}
