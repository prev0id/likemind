package credentials_adapter

import (
	"likemind/internal/database/model"
	"likemind/internal/domain"
)

func domainCredsToModel(creds domain.Credentials) model.Credentials {
	return model.Credentials{
		ID:       creds.ID,
		UserID:   creds.UserID,
		Password: creds.Password,
		Login:    creds.Login,
	}
}

func modelCredsToDomain(mCreds model.Credentials) domain.Credentials {
	return domain.Credentials{
		ID:       mCreds.ID,
		UserID:   mCreds.UserID,
		Password: mCreds.Password,
		Login:    mCreds.Login,
	}
}
