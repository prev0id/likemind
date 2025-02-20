package service

import (
	"context"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"slices"

	"likemind/internal/database/adapter/profile_adapter"
	"likemind/internal/domain"
	s3_image_repo "likemind/internal/s3/image_repo"

	"github.com/google/uuid"
)

const (
	MaxFileSize = 4 * 1024 * 1024

	MaxDimension = 2048

	MinAspectRatio = 8.0 / 19.0
	MaxAspectRatio = 19.0 / 8.0
)

type ImageService struct {
	s3   s3_image_repo.ImageRepository
	repo profile_adapter.Adapter
}

func NewImageService(
	s3 s3_image_repo.ImageRepository,
	repo profile_adapter.Adapter,
) *ImageService {
	return &ImageService{
		repo: repo,
		s3:   s3,
	}
}

func (s *ImageService) UploadImage(ctx context.Context, file io.Reader, fileSize int64) (string, error) {
	if fileSize > MaxFileSize {
		return "", domain.ErrFileSizeExceedsLimit
	}

	config, format, err := image.DecodeConfig(file)
	if err != nil {
		return "", domain.ErrInvalidFile
	}

	var contentType string
	switch format {
	case "jpeg":
		contentType = "image/jpeg"
	case "png":
		contentType = "image/png"
	default:
		return "", domain.ErrUsupportedImageFormat
	}

	if config.Width > MaxDimension || config.Height > MaxDimension {
		return "", domain.ErrWrongResolution
	}

	ratio := float64(config.Width) / float64(config.Height)
	if ratio < MinAspectRatio || ratio > MaxAspectRatio {
		return "", domain.ErrWrongAspectRation
	}

	id := uuid.New()
	var ext string
	if format == "jpeg" {
		ext = ".jpg"
	} else {
		ext = ".png"
	}
	uniqueFilename := id.String() + ext

	req := s3_image_repo.Image{
		Name:        uniqueFilename,
		ContentType: contentType,
		Size:        fileSize,
		Data:        file,
	}

	if err := s.s3.UploadImage(ctx, req); err != nil {
		return "", err
	}
	return uniqueFilename, nil
}

func (s *ImageService) DeleteImage(ctx context.Context, image string, userID int64) error {
	pics, err := s.repo.GetProfilePicturesByUserID(ctx, userID)
	if err != nil {
		return err
	}

	if !slices.Contains(pics, image) {
		return domain.ErrInvalidImageNameProvided
	}

	if err := s.repo.RemovePictureByID(ctx, image); err != nil {
		return err
	}

	if err := s.s3.DeleteImage(ctx, image); err != nil {
		return err
	}

	return nil
}

func (s *ImageService) GetImage(ctx context.Context, image string, userID int64) (io.ReadCloser, error) {
	pics, err := s.repo.GetProfilePicturesByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if !slices.Contains(pics, image) {
		return nil, domain.ErrInvalidImageNameProvided
	}

	data, err := s.s3.GetImage(ctx, image)
	if err != nil {
		return nil, err
	}

	return data, nil
}
