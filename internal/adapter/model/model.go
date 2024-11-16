package model

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:user"`
	ID            int64    `bun:"id"`
	Name          string   `bun:"name"`
	Contacts      []string `bun:"contacts,array"`
	Nickname      string   `bun:"nickname"`
	Surname       string   `bun:"surname"`
	PfpID         string   `bun:"pfp_id"`
	About         string   `bun:"about"`
}

// TODO:
type Interest struct{}

// TODO:
type UserInterest struct{}
