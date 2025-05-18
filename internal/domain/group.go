package domain

import (
	"strconv"
	"time"
)

type (
	GroupID   int64
	PostID    int64
	CommentID int64
)

type Group struct {
	ID          GroupID
	Name        string
	Description string
	Author      UserID
	Subscribed  bool
}

type Post struct {
	ID        PostID
	Author    UserID
	Group     GroupID
	Content   string
	Comments  []Comment
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Comment struct {
	ID        CommentID
	PostID    PostID
	Author    UserID
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (id GroupID) String() string {
	return strconv.FormatInt(int64(id), 10)
}

func (id PostID) String() string {
	return strconv.FormatInt(int64(id), 10)
}

func (id CommentID) String() string {
	return strconv.FormatInt(int64(id), 10)
}
