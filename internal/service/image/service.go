package image

import (
	"context"
	"fmt"
	"io"
	"slices"
	"strings"

	"likemind/internal/common"
	"likemind/internal/database/adapter/profile_adapter"
	"likemind/internal/domain"
	s3_image_repo "likemind/internal/s3/image_repo"

	"github.com/google/uuid"
	"github.com/ogen-go/ogen/http"
)

const (
	MaxFileSize = 4 * 1024 * 1024 // 4MB
)

var supportedFormats = []string{"image/png", "image/jpeg"}

type Service interface {
	UploadImage(ctx context.Context, file http.MultipartFile) error
	DeleteImage(ctx context.Context, image domain.PictureID, userID domain.UserID) error
	GetProfileImages(ctx context.Context, id domain.UserID) ([]domain.PictureID, error)
	GetImage(ctx context.Context, image domain.PictureID) (io.ReadCloser, error)
}

var _ (Service) = (*implementation)(nil)

type implementation struct {
	s3   s3_image_repo.ImageRepository
	repo profile_adapter.Adapter
}

func NewImageService(
	s3 s3_image_repo.ImageRepository,
	repo profile_adapter.Adapter,
) *implementation {
	return &implementation{
		repo: repo,
		s3:   s3,
	}
}

func (s *implementation) UploadImage(ctx context.Context, file http.MultipartFile) error {
	if file.File == nil {
		return fmt.Errorf("file reader is required")
	}
	if file.Size <= 0 {
		return fmt.Errorf("invalid file size")
	}
	if file.Size > MaxFileSize {
		return domain.ErrFileSizeExceedsLimit
	}

	contentType := file.Header.Get("Content-Type")

	if !slices.Contains(supportedFormats, contentType) {
		return domain.ErrUnsupportedImageFormat
	}

	id := uuid.New()
	uniqueFilename := id.String() + "." + strings.TrimPrefix(contentType, "image/")

	req := s3_image_repo.Image{
		Name:        uniqueFilename,
		ContentType: contentType,
		Size:        file.Size,
		Data:        file.File,
	}

	if err := s.s3.UploadImage(ctx, req); err != nil {
		return fmt.Errorf("failed to upload image: %w", err)
	}

	err := s.repo.AddProfilePicture(
		ctx,
		common.UserIDFromContext(ctx),
		domain.PictureID(uniqueFilename),
	)
	if err != nil {
		return fmt.Errorf("s.repo.AddProfilePicture: %w", err)
	}

	return nil
}

func (s *implementation) DeleteImage(ctx context.Context, image domain.PictureID, userID domain.UserID) error {
	pics, err := s.repo.GetProfilePicturesByUserID(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get profile pictures: %w", err)
	}

	if !slices.Contains(pics, image) {
		return domain.ErrInvalidImageNameProvided
	}

	if err := s.repo.RemovePictureByID(ctx, image); err != nil {
		return fmt.Errorf("failed to remove picture from DB: %w", err)
	}

	if err := s.s3.DeleteImage(ctx, string(image)); err != nil {
		return fmt.Errorf("failed to delete image from S3: %w", err)
	}

	return nil
}

func (s *implementation) GetImage(ctx context.Context, image domain.PictureID) (io.ReadCloser, error) {
	data, err := s.s3.GetImage(ctx, string(image))
	if err != nil {
		return nil, fmt.Errorf("failed to get image from S3: %w", err)
	}

	return data, nil
}

func (s *implementation) GetProfileImages(ctx context.Context, id domain.UserID) ([]domain.PictureID, error) {
	pictures, err := s.repo.GetProfilePicturesByUserID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("s.repo.GetProfilePicturesByUserID: %w", err)
	}
	return pictures, nil
}
