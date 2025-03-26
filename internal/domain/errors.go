package domain

import (
	"errors"
)

var (
	ErrNotAuthenticated         = errors.New("not authenticated")
	ErrUnauthorized             = errors.New("unauthorized")
	ErrInvalidImageNameProvided = errors.New("image not found")
	ErrFileSizeExceedsLimit     = errors.New("file size exceeds 4 MB")
	ErrInvalidFile              = errors.New("invalid file")
	ErrUnsupportedImageFormat   = errors.New("unsupported image format (only JPEG and PNG allowed)")
	ErrWrongResolution          = errors.New("image resolution exceeds 2K")
	ErrWrongAspectRation        = errors.New("image aspect ratio not allowed")
	ErrNilRequest               = errors.New("nil request")
	ErrNotFound                 = errors.New("not found")
)
