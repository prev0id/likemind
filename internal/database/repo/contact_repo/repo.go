package contact_repo

import (
	"context"
	"time"

	"likemind/internal/database"
	"likemind/internal/database/model"

	sql "github.com/huandu/go-sqlbuilder"
)

type DB interface {
	AddContact(ctx context.Context, contact model.Contact) error
	GetContactsByUserID(ctx context.Context, userID int64) ([]model.Contact, error)
	UpdateContact(ctx context.Context, picture model.Contact) error
	RemoveContactByID(ctx context.Context, id int64) error
}

var _ DB = (*Repo)(nil)

type Repo struct{}

func (r *Repo) AddContact(ctx context.Context, contact model.Contact) error {
	now := time.Now()
	contact.CreatedAt = now
	contact.UpdatedAt = now

	q := sql.InsertInto(model.TableContacts)
	q.Cols(
		model.ContactUserID,
		model.ContactPlatform,
		model.ContactValue,
		model.ContactCreatedAt,
		model.ContactUpdatedAt,
	)
	q.Values(
		contact.UserID,
		contact.Platform,
		contact.Value,
		contact.CreatedAt,
		contact.UpdatedAt,
	)

	if _, err := database.Exec(ctx, q); err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetContactsByUserID(ctx context.Context, userID int64) ([]model.Contact, error) {
	q := sql.Select(
		model.ContactID,
		model.ContactUserID,
		model.ContactPlatform,
		model.ContactValue,
		model.ContactCreatedAt,
		model.ContactUpdatedAt,
	)
	q.From(model.TableContacts)
	q.Where(q.Equal(model.ContactUserID, userID))

	results, err := database.Select[model.Contact](ctx, q)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r *Repo) UpdateContact(ctx context.Context, contact model.Contact) error {
	contact.UpdatedAt = time.Now()

	q := sql.Update(model.TableContacts)
	q.Set(
		q.Assign(model.ContactPlatform, contact.Platform),
		q.Assign(model.ContactValue, contact.Value),
		q.Assign(model.ContactUpdatedAt, contact.UpdatedAt),
	)
	q.Where(q.Equal(model.ContactID, contact.ID))

	if _, err := database.Exec(ctx, q); err != nil {
		return err
	}

	return nil
}

func (r *Repo) RemoveContactByID(ctx context.Context, id int64) error {
	q := sql.DeleteFrom(model.TableContacts)
	q.Where(q.Equal(model.ContactID, id))

	if _, err := database.Exec(ctx, q); err != nil {
		return err
	}

	return nil
}
