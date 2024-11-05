package domain

type User struct {
	ID       uint64
	Username string
	Name     string
	About    string
	Contacts []string
}
