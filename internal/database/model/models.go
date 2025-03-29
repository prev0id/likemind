package model

import "time"

const (
	TableSessions       = "sessions"
	TableUser           = "user"
	TableContact        = "contact"
	TableProfilePicture = "profile_picture"
	TableGroup          = "group"
	TableInterest       = "interest"
	TableUserInterest   = "user_interest"
	TableGroupInterest  = "group_interest"
)

const (
	CredentialsToken      = "token"
	CredentialsUserID     = "user_id"
	CredentialsCreatedAt  = "created_at"
	CredentialsExpectedAt = "expected_at"
)

type Session struct {
	UserID    int64     `db:"user_id"`
	Token     string    `db:"token"`
	CreatedAt time.Time `db:"created_at"`
	ExpiresAt time.Time `db:"expires_at"`
}

const (
	UserID        = "id"
	UserNickname  = "nickname"
	UserName      = "name"
	UserSurname   = "surname"
	UserAbout     = "about"
	UserPassword  = "password"
	UserEmail     = "email"
	UserLocation  = "location"
	UserCreatedAt = "created_at"
	UserUpdatedAt = "updated_at"
)

type User struct {
	ID        int64     `db:"id"`
	Nickname  string    `db:"nickname"`
	Name      string    `db:"name"`
	Surname   string    `db:"surname"`
	About     string    `db:"abuot"`
	Email     string    `db:"email"`
	Location  string    `db:"location"`
	Password  []byte    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

const (
	ContactID        = "id"
	ContactUserID    = "user_id"
	ContactPlatform  = "platform"
	ContactValue     = "contact"
	ContactCreatedAt = "created_at"
	ContactUpdatedAt = "updated_at"
)

type Contact struct {
	ID        int64     `db:"id"`
	UserID    int64     `db:"user_id"`
	Platform  string    `db:"platform"`
	Value     string    `db:"value"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

const (
	ProfilePictureID        = "id"
	ProfilePictureUserID    = "user_id"
	ProfilePictureCreatedAt = "created_at"
	ProfilePictureUpdatedAt = "updated_at"
)

type ProfilePicture struct {
	ID        string    `db:"id"`
	UserID    int64     `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

const (
	GroupID          = "id"
	GroupPictureID   = "picture_id"
	GroupName        = "name"
	GroupAlias       = "alias"
	GroupDescription = "description"
	GroupAuthorID    = "author_id"
	GroupCreatedAt   = "created_at"
	GroupUpdatedAt   = "updated_at"
)

type Group struct {
	ID          int64     `db:"id"`
	PictureID   string    `db:"picture_id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Alias       string    `db:"alias"`
	AuthorID    int64     `db:"author_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

const (
	InterestID          = "id"
	InterestName        = "name"
	InterestDescription = "description"
	InterestCreatedAt   = "created_at"
	InterestUpdatedAt   = "updated_at"
)

type Interest struct {
	ID          int64     `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

const (
	UserInterestUserID     = "user_id"
	UserInterestInterestID = "name"
	UserInterestCreatedAt  = "created_at"
)

type UserInterest struct {
	UserID     int64     `db:"user_id"`
	InterestID int64     `db:"interest_id"`
	CreatedAt  time.Time `db:"created_at"`
}

const (
	GroupInterestGroupID    = "group_id"
	GroupInterestInterestID = "name"
	GroupInterestCreatedAt  = "created_at"
)

type GroupInterest struct {
	GroupID    int64     `db:"group_id"`
	InterestID int64     `db:"interest_id"`
	CreatedAt  time.Time `db:"created_at"`
}
