package validator

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const (
	NotBlankField = "this field cannot be blank"
	ErrMaxLength  = "cannot be longer than %d characters"
	ErrMinLength  = "must be at least %d characters"
)

type Validator struct {
	FieldErrors map[string]string
}

func NewValidator() *Validator {
	return &Validator{
		FieldErrors: make(map[string]string),
	}
}

func (v *Validator) HasErrors() bool {
	return len(v.FieldErrors) > 0
}

func (v *Validator) AddFieldError(key, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}

	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}
}

func (v *Validator) Assert(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

func (v *Validator) Errors() []string {
	errs := make([]string, 0, len(v.FieldErrors))

	for key, value := range v.FieldErrors {
		errs = append(errs, fmt.Sprintf(`%s: %s`, key, value))
	}
	return errs
}

func NotBlank(value string) bool {
	// Trim splaces to guard for "\t\t\t\n\n\n\    ".
	return strings.TrimSpace(value) != ""
}

// MaxChars & MinChars checks for utf8 chars.

func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

func MinChars(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n
}

func (v *Validator) Merge(other *Validator) {
	if other == nil {
		return
	}

	for key, message := range other.FieldErrors {
		v.AddFieldError(key, message)
	}
}
