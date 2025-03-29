package domain

type Interest struct {
	ID          int64
	Name        string
	Description string
	GroupID     int64
}

type InterestGroup struct {
	ID   int64
	Name string
}
