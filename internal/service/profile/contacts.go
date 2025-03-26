package profile

import (
	"context"
	"fmt"
	"slices"

	"likemind/internal/domain"
)

const idxNotFound = -1

func (s *implementation) AddContact(ctx context.Context, id domain.UserID, contact domain.Contact) error {
	if err := s.db.AddContact(ctx, id, contact); err != nil {
		return fmt.Errorf("s.db.AddContact: %w", err)
	}
	return nil
}

func (s *implementation) UpdateContact(ctx context.Context, id domain.UserID, contact domain.Contact) error {
	if err := s.validateContactOwnership(ctx, id, contact.ID); err != nil {
		return err
	}

	if err := s.db.UpdateContact(ctx, id, contact); err != nil {
		return fmt.Errorf("s.db.UpdateContact: %w", err)
	}
	return nil
}

func (s *implementation) RemoveContact(ctx context.Context, id domain.UserID, contactID int64) error {
	if err := s.validateContactOwnership(ctx, id, contactID); err != nil {
		return err
	}

	if err := s.db.RemoveContactByID(ctx, contactID); err != nil {
		return fmt.Errorf("s.db.RemoveContactByID: %w", err)
	}
	return nil
}

func (s *implementation) validateContactOwnership(ctx context.Context, userID domain.UserID, contactID int64) error {
	contacts, err := s.db.GetContactsByUserID(ctx, userID)
	if err != nil {
		return fmt.Errorf("s.db.GetContactByID: %w", err)
	}

	idx := slices.IndexFunc(contacts, func(contact domain.Contact) bool {
		return contact.ID == contactID
	})

	if idx != idxNotFound {
		return domain.ErrNotFound
	}

	return nil
}
