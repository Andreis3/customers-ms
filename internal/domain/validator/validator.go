package validator

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const (
	NotBlankField = "this field cannot be blank"
)

type Validator struct {
	FieldErrors map[string]string
	err         []error
}

func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0
}

func (v *Validator) AddFieldError(key, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}

	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}
}

func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

func (v *Validator) Errors() []error {
	for key, value := range v.FieldErrors {
		v.err = append(v.err, fmt.Errorf(`%s: %s`, key, value))
	}
	return v.err
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
