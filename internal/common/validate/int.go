package validate

import "fmt"

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type IntValidator[T Integer] struct {
	FieldName   string
	rules       []func(T) error
	failOnFirst bool
}

func Int[T Integer](fieldName string) *IntValidator[T] {
	return &IntValidator[T]{
		FieldName: fieldName,
		rules:     []func(T) error{},
	}
}

func (iv *IntValidator[T]) NotEmpty() *IntValidator[T] {
	rule := func(val T) error {
		var zero T
		if val == zero {
			return fmt.Errorf("%s should not be empty: %w", iv.FieldName, ErrNotEmpty)
		}
		return nil
	}
	iv.rules = append(iv.rules, rule)
	return iv
}

func (iv *IntValidator[T]) Min(min T) *IntValidator[T] {
	rule := func(val T) error {
		if val < min {
			return fmt.Errorf("%s must be at least %v: %w", iv.FieldName, min, ErrMin)
		}
		return nil
	}
	iv.rules = append(iv.rules, rule)
	return iv
}

func (iv *IntValidator[T]) Max(max T) *IntValidator[T] {
	rule := func(val T) error {
		if val > max {
			return fmt.Errorf("%s must be at most %v: %w", iv.FieldName, max, ErrMax)
		}
		return nil
	}
	iv.rules = append(iv.rules, rule)
	return iv
}

func (iv *IntValidator[T]) FailOnFirst() *IntValidator[T] {
	iv.failOnFirst = true
	return iv
}

func (iv *IntValidator[T]) Build(value T) error {
	if iv.failOnFirst {
		return validate(value, iv.rules)
	}

	return validateAll(value, iv.rules)
}
