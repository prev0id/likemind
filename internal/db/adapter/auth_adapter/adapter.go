package auth_adapter

import (
	"context"

	"likemind/internal/db/data_provider"
	"likemind/internal/db/model"
	"likemind/internal/db/op"
	"likemind/internal/domain"
)

type Adpater struct {
	provider data_provider.DataProvider[domain.Credential, model.Credential, string]
}

func (a *Adpater) Create(ctx context.Context, creds domain.Credential) error {
	_, err := a.provider.
		Insert(ctx).
		Field(model.CredentialUuid, creds.UUID).
		Field(model.CredentialPassword, creds.Password).
		Field(model.CredentialLogin, creds.Login).
		Field(model.CredentialUserId, creds.UserID).
		Field(model.CredentialCreatedAt, creds.CreatedAt).
		Field(model.CredentialUpdatedAt, creds.UpdatedAt).
		Do(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (a *Adpater) GetByLogin(ctx context.Context, login string) (domain.Credential, error) {
	creds, err := a.provider.
		Get(ctx).
		ByFilter(model.CredentialLogin, op.Eq, login).
		Do(ctx)
	if err != nil {
		return domain.Credential{}, err
	}

	return creds, nil
}

func (a *Adpater) GetByUUID(ctx context.Context, uuid string) (domain.Credential, error) {
	creds, err := a.provider.
		Get(ctx).
		ByPK(uuid).
		Do(ctx)
	if err != nil {
		return domain.Credential{}, err
	}

	return creds, nil
}
