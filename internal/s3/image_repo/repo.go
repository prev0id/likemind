package s3_image_repo

import (
	"context"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type ImageRepository interface {
	UploadImage(ctx context.Context, image Image) error
	GetImage(ctx context.Context, name string) (io.ReadCloser, error)
	DeleteImage(ctx context.Context, name string) error
}

type S3Repository struct {
	client *minio.Client
	cfg    config
}

type Image struct {
	Name        string
	ContentType string
	Data        io.Reader
	Size        int64
}

type config struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	BucketName      string
	Location        string
	UseSSL          bool
}

func NewS3Repository(cfg config) (ImageRepository, error) {
	minioClient, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		Secure: cfg.UseSSL,
	})
	if err != nil {
		return nil, err
	}

	return &S3Repository{
		client: minioClient,
		cfg:    cfg,
	}, nil
}

func (r *S3Repository) UploadImage(ctx context.Context, image Image) error {
	_, err := r.client.PutObject(ctx, r.cfg.BucketName, image.Name, image.Data, image.Size, minio.PutObjectOptions{ContentType: image.ContentType})
	if err != nil {
		return err
	}
	return nil
}

func (r *S3Repository) GetImage(ctx context.Context, name string) (io.ReadCloser, error) {
	reader, err := r.client.GetObject(context.Background(), r.cfg.BucketName, name, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	return reader, nil
}

func (r *S3Repository) DeleteImage(ctx context.Context, name string) error {
	return r.client.RemoveObject(context.Background(), r.cfg.BucketName, name, minio.RemoveObjectOptions{})
}
