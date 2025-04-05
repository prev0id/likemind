package profile

import (
	"context"
	"fmt"

	"likemind/internal/domain"
)

func (s *implementation) AddProfilePicture(ctx context.Context, id domain.UserID, pictureID domain.PictureID) error {
	if err := s.db.AddProfilePicture(ctx, id, pictureID); err != nil {
		return fmt.Errorf("s.db.AddProfilePicture: %w", err)
	}
	return nil
}

func (s *implementation) GetProfilePictures(ctx context.Context, id domain.UserID) ([]domain.PictureID, error) {
	pictures, err := s.db.GetProfilePicturesByUserID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("s.db.GetProfilePicturesByUserID: %w", err)
	}
	return pictures, nil
}

func (s *implementation) RemovePicture(ctx context.Context, pictureID domain.PictureID) error {
	if err := s.db.RemovePictureByID(ctx, pictureID); err != nil {
		return fmt.Errorf("s.db.RemovePictureByID: %w", err)
	}
	return nil
}
