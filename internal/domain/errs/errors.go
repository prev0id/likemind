package errs

import "net/http"

func InternalError(internalMessage string) *Error {
	return &Error{
		StatusCode: http.StatusInternalServerError,
		Internal:   internalMessage,
		Public:     "Same unexpected error occurred, try request later",
	}
}

func PermissionDeniedError(permissionMessage string) *Error {
	return &Error{
		StatusCode: http.StatusForbidden,
		Internal:   permissionMessage,
		Public:     "Permission denied, you should not do this",
	}
}

func NotAuthenticatedError(authMessage string) *Error {
	return &Error{
		StatusCode: http.StatusUnauthorized,
		Internal:   authMessage,
		Public:     "Authentication is required, please provide valid credentials",
	}
}

func UserNotFoundError(userMessage string) *Error {
	return &Error{
		StatusCode: http.StatusNotFound,
		Internal:   userMessage,
		Public:     "User not found, please check the user identifier and try again",
	}
}

func GroupNotFoundError(groupMessage string) *Error {
	return &Error{
		StatusCode: http.StatusNotFound,
		Internal:   groupMessage,
		Public:     "Group not found, please check the group identifier and try again",
	}
}
