package common

import (
	"errors"
	"io"
	"strings"

	"likemind/internal/domain"
)

type errorType int

const (
	InternalErrorType errorType = iota
	NotFoundErrorType
	BadRequestErrorType
	NotAuthenticatedErrorType
	PermissionDeniedErrorType
)

var domainErrors = map[error]errorType{
	domain.ErrNotAuthenticated:         NotAuthenticatedErrorType,
	domain.ErrUnauthorized:             PermissionDeniedErrorType,
	domain.ErrInvalidImageNameProvided: BadRequestErrorType,
	domain.ErrFileSizeExceedsLimit:     BadRequestErrorType,
	domain.ErrInvalidFile:              BadRequestErrorType,
	domain.ErrUnsupportedImageFormat:   BadRequestErrorType,
	domain.ErrWrongResolution:          BadRequestErrorType,
	domain.ErrWrongAspectRation:        BadRequestErrorType,
	domain.ErrValidationFailed:         BadRequestErrorType,
}

func ErrorMsg(err error) io.Reader {
	if err == nil {
		return nil
	}
	return strings.NewReader(err.Error())
}

func ErrorIs(err error, errType errorType) bool {
	for domainErr, domainErrType := range domainErrors {
		if errors.Is(err, domainErr) {
			return errType == domainErrType
		}
	}
	return false
}
