package view

import "time"

type Group struct {
	ID             int64
	Name           string
	Description    string
	AuthorNickname string
	Posts          []Post
	Interests      []Interest
}

type Post struct {
	Content        string
	AuthorNickname string
	Comments       []Comment
	CreatedAt      time.Time
}

type Comment struct {
	Content        string
	AuthorNickname string
	CreatedAt      time.Time
}
