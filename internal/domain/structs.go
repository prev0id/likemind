package domain

const (
	TableUser          = "user"
	TableGroup         = "group"
	TableInterest      = "interest"
	TableCredential    = "credential"
	TableUserInterest  = "user_interest"
	TableGroupInterest = "group_interest"

	FieldID       = "id"
	FieldUUID     = "uuid"
	FieldNickname = "nickname"
	FieldLogin    = "login"
)

type User struct {
	DataFields
	Nickname string   `db:"nickname"`
	Name     string   `db:"name"`
	Surname  string   `db:"surname"`
	About    string   `db:"about"`
	PfpID    string   `db:"pfp_id"`
	Contacts []string `db:"contacts"`
}

type Interest struct {
	DataFields
	Name        string `db:"name"`
	Description string `db:"description"`
}

type Group struct {
	DataFields
	Name        string `db:"name"`
	Description string `db:"description"`
	PictureID   string `db:"picture_id"`
}

type Post struct {
	DataFields
	GroupID  int64  `db:"group_id"`
	AuthorID int64  `db:"author_id"`
	Text     string `db:"text"`
}

type AppliedInterest struct {
	DataFields
	InterestID int64 `db:"interest_id"`
	EntityID   int64 `db:"entity_id"`
}

type Credential struct {
	DataFields
	Password []byte `db:"password"`
	Login    string `db:"login"`
	UserID   int64  `db:"user_id"`
	UUID     string `db:"uuid"`
}
