package errors

import "github.com/andreis3/users-ms/internal/domain/validator"

type ErrorCode string

const (
	ErrInvalidBusinessRules ErrorCode = "APP-400"
	ErrResourceNotFound     ErrorCode = "APP-404"
	ErrInternalProcessing   ErrorCode = "APP-500"
)

type DomainError struct {
	Code            ErrorCode
	Message         []string
	OriginFunc      string
	FriendlyMessage string
}

func (de DomainError) Error() string {
	return string(de.Code) + ": " + de.FriendlyMessage
}

func InvalidCustomerError(validate *validator.Validator) *DomainError {
	return &DomainError{
		Code:            ErrInvalidBusinessRules,
		Message:         validate.Errors(),
		OriginFunc:      "CustomerProfile.Validate",
		FriendlyMessage: "Validation failed for the provided input.",
	}
}
