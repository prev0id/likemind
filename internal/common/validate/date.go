package validate

import (
	"fmt"
	"time"
)

type DateValidator struct {
	FieldName   string
	rules       []func(time.Time) error
	failOnFirst bool
}

func Date(fieldName string) *DateValidator {
	return &DateValidator{
		FieldName: fieldName,
		rules:     []func(time.Time) error{},
	}
}

func (dv *DateValidator) NotEmpty() *DateValidator {
	rule := func(val time.Time) error {
		if val.IsZero() {
			return fmt.Errorf("%s should not be empty: %w", dv.FieldName, ErrNotEmpty)
		}
		return nil
	}
	dv.rules = append(dv.rules, rule)
	return dv
}

func (dv *DateValidator) Before(t time.Time) *DateValidator {
	rule := func(val time.Time) error {
		if !val.Before(t) {
			return fmt.Errorf("%s must be before %s: %w", dv.FieldName, t.Format(time.RFC3339), ErrPattern)
		}
		return nil
	}
	dv.rules = append(dv.rules, rule)
	return dv
}

func (dv *DateValidator) After(t time.Time) *DateValidator {
	rule := func(val time.Time) error {
		if !val.After(t) {
			return fmt.Errorf("%s must be after %s: %w", dv.FieldName, t.Format(time.RFC3339), ErrPattern)
		}
		return nil
	}
	dv.rules = append(dv.rules, rule)
	return dv
}

func (dv *DateValidator) FailOnFirst() *DateValidator {
	dv.failOnFirst = true
	return dv
}

func (dv *DateValidator) Build(value time.Time) error {
	if dv.failOnFirst {
		return validate(value, dv.rules)
	}

	return validateAll(value, dv.rules)
}
