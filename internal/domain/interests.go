package domain

type Interests []InterestGroup

type InterestGroup struct {
	ID        int64
	Name      string
	Interests []Interest
}

type Interest struct {
	ID          int64
	Name        string
	Description string
	GroupID     int64
	Selected    bool
}
