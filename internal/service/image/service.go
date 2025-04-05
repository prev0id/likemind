package service

import (
	"context"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"slices"
	"strings"

	"likemind/internal/database/adapter/profile_adapter"
	"likemind/internal/domain"
	s3_image_repo "likemind/internal/s3/image_repo"

	"github.com/google/uuid"
)

const (
	MaxFileSize = 4 * 1024 * 1024 // 4MB

	MaxDimension = 2048

	MinAspectRatio = 8.0 / 19.0
	MaxAspectRatio = 19.0 / 8.0

	supportedFormats = "jpeg,png"
)

type ImageService struct {
	s3   s3_image_repo.ImageRepository
	repo profile_adapter.Adapter
}

func NewImageService(
	s3 s3_image_repo.ImageRepository,
	repo profile_adapter.Adapter,
) (*ImageService, error) {
	if s3 == nil {
		return nil, fmt.Errorf("s3 repository is required")
	}
	if repo == nil {
		return nil, fmt.Errorf("profile adapter is required")
	}

	return &ImageService{
		repo: repo,
		s3:   s3,
	}, nil
}

func (s *ImageService) UploadImage(ctx context.Context, file io.Reader, fileSize int64) (string, error) {
	if file == nil {
		return "", fmt.Errorf("file reader is required")
	}
	if fileSize <= 0 {
		return "", fmt.Errorf("invalid file size")
	}
	if fileSize > MaxFileSize {
		return "", domain.ErrFileSizeExceedsLimit
	}

	config, format, err := image.DecodeConfig(file)
	if err != nil {
		return "", fmt.Errorf("failed to decode image: %w", domain.ErrInvalidFile)
	}

	format = strings.ToLower(format)
	if !strings.Contains(supportedFormats, format) {
		return "", domain.ErrUnsupportedImageFormat
	}

	contentType := fmt.Sprintf("image/%s", format)

	if config.Width > MaxDimension || config.Height > MaxDimension {
		return "", domain.ErrWrongResolution
	}

	ratio := float64(config.Width) / float64(config.Height)
	if ratio < MinAspectRatio || ratio > MaxAspectRatio {
		return "", domain.ErrWrongAspectRation
	}

	id := uuid.New()
	ext := fmt.Sprintf(".%s", strings.ReplaceAll(format, "jpeg", "jpg"))
	uniqueFilename := id.String() + ext

	req := s3_image_repo.Image{
		Name:        uniqueFilename,
		ContentType: contentType,
		Size:        fileSize,
		Data:        file,
	}

	if err := s.s3.UploadImage(ctx, req); err != nil {
		return "", fmt.Errorf("failed to upload image: %w", err)
	}

	return uniqueFilename, nil
}

func (s *ImageService) DeleteImage(ctx context.Context, image domain.PictureID, userID domain.UserID) error {
	if image == "" {
		return fmt.Errorf("image name is required")
	}
	if userID <= 0 {
		return fmt.Errorf("invalid user ID")
	}

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

func (s *ImageService) GetImage(ctx context.Context, image domain.PictureID, userID domain.UserID) (io.ReadCloser, error) {
	if ctx == nil {
		return nil, fmt.Errorf("context is required")
	}
	if image == "" {
		return nil, fmt.Errorf("image name is required")
	}
	if userID <= 0 {
		return nil, fmt.Errorf("invalid user ID")
	}

	pics, err := s.repo.GetProfilePicturesByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get profile pictures: %w", err)
	}

	if !slices.Contains(pics, image) {
		return nil, domain.ErrInvalidImageNameProvided
	}

	data, err := s.s3.GetImage(ctx, string(image))
	if err != nil {
		return nil, fmt.Errorf("failed to get image from S3: %w", err)
	}

	return data, nil
}
