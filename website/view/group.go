package view

import "time"

type Group struct {
	ID          int64
	Name        string
	Description string
	Author      *Profile
	Posts       []*Post
	Interests   []GroupedInterests
}

type Post struct {
	ID        int64
	Content   string
	Author    *Profile
	Comments  []*Comment
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Comment struct {
	ID        int64
	Content   string
	Author    *Profile
	CreatedAt time.Time
	UpdatedAt time.Time
}
