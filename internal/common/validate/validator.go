package validate

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	ErrValidation = errors.New("validation failed")
	ErrMin        = errors.New("value below minimum")
	ErrMax        = errors.New("value above maximum")
	ErrLen        = errors.New("invalid length")
	ErrPattern    = errors.New("pattern mismatch")
	ErrNotEmpty   = errors.New("value is empty")
)

var (
	PatternEmail = Pattern{
		Regex: regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`),
		Name:  "email",
	}
	PatternName = Pattern{
		Regex: regexp.MustCompile(`^[A-Za-z\s\-]+$`),
		Name:  "name",
	}
	Location = Pattern{
		Regex: regexp.MustCompile(`^[A-Za-z0-9\s,\-]+$`),
		Name:  "location",
	}
)

type Pattern struct {
	Regex *regexp.Regexp
	Name  string
}

func validate[T any](value T, rules []func(T) error) error {
	for _, rule := range rules {
		if err := rule(value); err != nil {
			return fmt.Errorf("%w: %w", ErrValidation, err)
		}
	}

	return nil
}

func validateAll[T any](value T, rules []func(T) error) error {
	var errs []string

	for _, rule := range rules {
		if err := rule(value); err != nil {
			errs = append(errs, err.Error())
		}
	}

	if len(errs) == 0 {
		return nil
	}

	result := strings.Builder{}
	for i, errMsg := range errs {
		result.WriteString(strconv.Itoa(i + 1))
		result.WriteString(") ")
		result.WriteString(errMsg)
		result.WriteRune('\n')
	}

	return fmt.Errorf("%w\n: %s", ErrValidation, result.String())
}
