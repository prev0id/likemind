package domain

import "time"

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
	DateOfBirth    time.Time
	HashedPassword []byte
	RawPassword    Password
}

type Contact struct {
	ID       int64
	Platform string
	Value    string
}
