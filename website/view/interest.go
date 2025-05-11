package view

type Interest struct {
	ID          int64
	Name        string
	Description string
	Selected    bool
}

type GroupedInterests struct {
	Name      string
	Interests []Interest
}
