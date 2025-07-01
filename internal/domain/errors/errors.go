package errors

import (
	"errors"
)

type Error struct {
	Code            Code
	Errors          []string
	Map             map[string]any
	OriginFunc      string
	Cause           string
	FriendlyMessage string
}

type InputError Error

func (e *Error) Error() string {
	return string(e.Code) + ": " + e.FriendlyMessage
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}

func Join(errs ...error) error {
	return errors.Join(errs...)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func New(inputError InputError) *Error {
	return &Error{
		Code:            inputError.Code,
		Errors:          inputError.Errors,
		OriginFunc:      inputError.OriginFunc,
		FriendlyMessage: inputError.FriendlyMessage,
	}
}
