package domain

const (
	TableUser          = "user"
	TableGroup         = "group"
	TableInterest      = "interest"
	TableCredential    = "credential"
	TableUserInterest  = "user_interest"
	TableGroupInterest = "group_interest"
)

type User struct {
	DataFields
	Nickname string
	Name     string
	Surname  string
	About    string
	PfpID    string
	Contacts []string
}

type Interest struct {
	DataFields
	Name        string
	Description string
}

type Group struct {
	DataFields
	Name        string
	Description string
	PictureID   string
}

type Post struct {
	DataFields
	GroupID  int64
	AuthorID int64
	Text     string
}

type AppliedInterest struct {
	DataFields
	InterestID int64
	EntityID   int64
}

type Credential struct {
	Data
	Password []byte
	Login    string
	UserID   int64
	UUID     string
}
