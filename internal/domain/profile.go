package domain

import "time"

type (
	UserID   int64
	Password string
	Email    string
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
	ID             UserID
	Nickname       string
	Name           string
	Surname        string
	About          string
	Login          Email
	DateOfBirth    time.Time
	HashedPassword []byte
	RawPassword    Password
}
