package validate

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

type StringValidator struct {
	FieldName   string
	failOnFirst bool
	rules       []func(string) error
}

func String(fieldName string) *StringValidator {
	return &StringValidator{
		FieldName: fieldName,
		rules:     []func(string) error{},
	}
}

func (sv *StringValidator) NotEmpty() *StringValidator {
	rule := func(val string) error {
		if val == "" {
			return fmt.Errorf("'%s' should not be empty: %w", sv.FieldName, ErrNotEmpty)
		}
		return nil
	}
	sv.rules = append(sv.rules, rule)
	return sv
}

func (sv *StringValidator) LenMin(min int) *StringValidator {
	rule := func(val string) error {
		if len(val) < min {
			return fmt.Errorf("'%s' must be at least %d characters: %w", sv.FieldName, min, ErrMin)
		}
		return nil
	}
	sv.rules = append(sv.rules, rule)
	return sv
}

func (sv *StringValidator) LenMax(max int) *StringValidator {
	rule := func(val string) error {
		if len(val) > max {
			return fmt.Errorf("'%s' must be at most %d characters: %w", sv.FieldName, max, ErrMax)
		}
		return nil
	}
	sv.rules = append(sv.rules, rule)
	return sv
}

func (sv *StringValidator) LenExect(length int) *StringValidator {
	rule := func(val string) error {
		if len(val) != length {
			return fmt.Errorf("'%s' must be exactly %d characters: %w", sv.FieldName, length, ErrLen)
		}
		return nil
	}
	sv.rules = append(sv.rules, rule)
	return sv
}

func (sv *StringValidator) Pattern(p Pattern) *StringValidator {
	rule := func(val string) error {
		if !p.Regex.MatchString(val) {
			return fmt.Errorf("'%s' does not match the '%s' pattern: %w", sv.FieldName, p.Name, ErrPattern)
		}
		return nil
	}
	sv.rules = append(sv.rules, rule)
	return sv
}

func (sv *StringValidator) ContainsUpper() *StringValidator {
	rule := func(val string) error {
		for _, r := range val {
			if unicode.IsUpper(r) {
				return nil
			}
		}
		return fmt.Errorf("'%s' must contain at least one uppercase letter: %w", sv.FieldName, ErrPattern)
	}
	sv.rules = append(sv.rules, rule)
	return sv
}

func (sv *StringValidator) ContainsLower() *StringValidator {
	rule := func(val string) error {
		for _, r := range val {
			if unicode.IsLower(r) {
				return nil
			}
		}
		return fmt.Errorf("'%s' must contain at least one lowercase letter: %w", sv.FieldName, ErrPattern)
	}
	sv.rules = append(sv.rules, rule)
	return sv
}

func (sv *StringValidator) ContainsDigit() *StringValidator {
	rule := func(val string) error {
		for _, r := range val {
			if unicode.IsDigit(r) {
				return nil
			}
		}
		return fmt.Errorf("'%s' must contain at least one digit: %w", sv.FieldName, ErrPattern)
	}
	sv.rules = append(sv.rules, rule)
	return sv
}

func (sv *StringValidator) ContainsSymbol() *StringValidator {
	rule := func(val string) error {
		for _, r := range val {
			if unicode.IsPunct(r) || unicode.IsSymbol(r) {
				return nil
			}
		}
		return fmt.Errorf("'%s' must contain at least one symbol: %w", sv.FieldName, ErrPattern)
	}
	sv.rules = append(sv.rules, rule)
	return sv
}

func (sv *StringValidator) IsUTF8() *StringValidator {
	rule := func(val string) error {
		if !utf8.ValidString(val) {
			return fmt.Errorf("'%s' is not a valid UTF-8 string: %w", sv.FieldName, ErrPattern)
		}
		return nil
	}
	sv.rules = append(sv.rules, rule)
	return sv
}

func (sv *StringValidator) FailOnFirst() *StringValidator {
	sv.failOnFirst = true
	return sv
}

func (sv *StringValidator) Build(value string) error {
	if sv.failOnFirst {
		return validate(value, sv.rules)
	}

	return validateAll(value, sv.rules)
}
