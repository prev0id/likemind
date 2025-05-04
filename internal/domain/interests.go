package domain

type Interest struct {
	ID          int64
	Name        string
	Description string
	GroupID     int64
}

type Interests map[InterestGroup]Interest

type InterestGroup struct {
	ID   int64
	Name string
}

type UserInterest struct {
	UserID   UserID
	Interest []Interest
}
