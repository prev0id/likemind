package common

import (
	"errors"
	"io"
	"likemind/internal/domain"
	"strings"
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
	domain.ErrUsupportedImageFormat:    BadRequestErrorType,
	domain.ErrWrongResolution:          BadRequestErrorType,
	domain.ErrWrongAspectRation:        BadRequestErrorType,
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
