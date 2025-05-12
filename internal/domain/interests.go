package domain

type InterestID int64

type InterestGroup struct {
	Name      string
	Interests []Interest
}

type Interest struct {
	ID          InterestID
	Name        string
	Description string
	GroupID     int64
	Selected    bool
}
