package view

import "time"

type Profile struct {
	Nickname    string
	Name        string
	Surname     string
	About       string
	Location    string
	DateOfBirth time.Time

	Contacts  []Contact
	PictureID string

	Interests []GroupedInterests

	Owner bool
}

type Contact struct {
	ID       int64
	Platform string
	Value    string
}
