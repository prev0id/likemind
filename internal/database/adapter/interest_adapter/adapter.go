package interest_adapter

import "likemind/internal/database/repo/interest_repo"

type Adapter interface{}

type Implementation struct {
	Repo *interest_repo.DB
}

func (i *Implementation) ListUsersInterests() {
}
