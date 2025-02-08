package model

//go:generate go-enum --marshal --values --names

// ENUM(user, group, interest, credential, user_interest, group_interest)
type Table int

// ENUM(id, nickname, name, surname, pfp_id, about, contacts, created_at, updated_at)
type User int

// ENUM(uuid, password, login, user_id, created_at, updated_at)
type Credential int

type M[Field F, PrimaryKey comparable] interface {
	TableName() Table
	FieldPrimaryKey() Field
	FieldUpdateAt() Field
	FieldCreatedAt() Field
}

type F interface {
	String() string
}
