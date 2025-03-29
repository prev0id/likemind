package domain

import "time"

type Group struct {
	Info      GroupInfo
	Posts     []Post
	Interests []Interest
}

type GroupInfo struct {
	ID          int64
	Name        string
	Description string
	Author      UserID
}

type Post struct {
	ID        int64
	Content   string
	Author    UserID
	Comments  []Comment
	CreatedAt time.Time
}

type Comment struct {
	ID        int64
	PostID    int64
	Content   string
	Author    UserID
	CreatedAt time.Time
}
