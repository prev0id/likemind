package credentials_adapter

import (
	"context"

	"likemind/internal/database/repo/credentials_repo"
	"likemind/internal/domain"
)

type Adapter interface {
	Create(ctx context.Context, creds domain.Credentials) error
	Update(ctx context.Context, creds domain.Credentials) error
	GetByLogin(ctx context.Context, login domain.Email) (domain.Credentials, error)
	GetByID(ctx context.Context, id string) (domain.Credentials, error)
}

type implementation struct {
	credsDB credentials_repo.DB
}

func NewAdapter(credsDB credentials_repo.DB) Adapter {
	return &implementation{
		credsDB: credsDB,
	}
}

func (i *implementation) Create(ctx context.Context, creds domain.Credentials) error {
	return i.credsDB.Create(ctx, domainCredsToModel(creds))
}

func (i *implementation) Update(ctx context.Context, creds domain.Credentials) error {
	return i.credsDB.Update(ctx, domainCredsToModel(creds))
}

func (i *implementation) GetByLogin(ctx context.Context, login domain.Email) (domain.Credentials, error) {
	creds, err := i.credsDB.GetByLogin(ctx, string(login))
	if err != nil {
		return domain.Credentials{}, err
	}
	return modelCredsToDomain(creds), nil
}

func (i *implementation) GetByID(ctx context.Context, id string) (domain.Credentials, error) {
	creds, err := i.credsDB.GetByID(ctx, id)
	if err != nil {
		return domain.Credentials{}, err
	}
	return modelCredsToDomain(creds), nil
}
