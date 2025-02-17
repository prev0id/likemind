package domain

import (
	"errors"
	"net/http"
)

var (
	ErrUnauthenticated          = errors.New("unauthenticated")
	ErrUnauthorized             = errors.New("unauthorized")
	ErrInvalidSession           = errors.New("invalid session")
	ErrNoSession                = errors.New("no session cookie")
	ErrInvalidImageNameProvided = errors.New("image not found")
	ErrFileSizeExceedsLimit     = errors.New("file size exceeds 4 MB")
	ErrInvalidFile              = errors.New("invalid file")
	ErrUsupportedImageFormat    = errors.New("unsupported image format (only JPEG and PNG allowed)")
	ErrWrongResolution          = errors.New("image resolution exceeds 2K")
	ErrWrongAspectRation        = errors.New("image aspect ratio not allowed")
)

var errToStatusCode = map[error]int{
	ErrUnauthorized:             http.StatusUnauthorized,
	ErrUnauthenticated:          http.StatusForbidden,
	ErrInvalidImageNameProvided: http.StatusUnauthorized,
	ErrFileSizeExceedsLimit:     http.StatusBadRequest,
	ErrInvalidFile:              http.StatusBadRequest,
	ErrUsupportedImageFormat:    http.StatusBadRequest,
	ErrWrongResolution:          http.StatusBadRequest,
	ErrWrongAspectRation:        http.StatusBadRequest,
}
