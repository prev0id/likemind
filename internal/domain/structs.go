package domain

type User struct {
	Username string
	Name     string
	Surname  string
	About    string
	Contacts []string
	ID       int64
	PfpID    string
}

type AppliedInterest struct {
	UserID     uint64
	InterestID uint64
	IsLiked    bool
}

type Interest struct {
	Name        string
	Description string
	ID          uint64
}
