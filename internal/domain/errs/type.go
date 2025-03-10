package errs

import (
	"fmt"
	"strconv"
	"strings"
)

type Error struct {
	StatusCode int
	Internal   string
	Public     string
	Stack      []string
}

func (e *Error) Error() string {
	return e.Public
}

func Stack(err error, format string, args ...any) error {
	if err == nil {
		return nil
	}

	message := fmt.Sprintf(format, args...)

	internalError, ok := err.(*Error)
	if !ok {
		internalError = InternalError(err.Error())
	}

	internalError.Stack = append(internalError.Stack, message)

	return internalError
}

func (e *Error) PrintInternalError() string {
	if e == nil {
		return ""
	}

	var builder strings.Builder

	if e.StatusCode != 0 {
		builder.WriteString("[status_code] ")
		builder.WriteString(strconv.Itoa(e.StatusCode))
		builder.WriteRune('\n')
	}

	if e.Public != "" {
		builder.WriteString("[public] ")
		builder.WriteString(e.Public)
		builder.WriteRune('\n')
	}

	if e.Internal != "" {
		builder.WriteString("[internal] ")
		builder.WriteString(e.Internal)
		builder.WriteRune('\n')
	}

	for _, stackLine := range e.Stack {
		builder.WriteString("->")
		builder.WriteString(stackLine)
		builder.WriteRune('\n')
	}

	return builder.String()
}
