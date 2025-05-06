package apperrors

import "github.com/andreis3/customers-ms/internal/domain/validator"

type ErrorCode string

const (
	ErrInvalidBusinessRules ErrorCode = "DM-400"
	ErrResourceNotFound     ErrorCode = "DM-404"
)

func InvalidCustomerError(validate *validator.Validator) *AppErrors {
	return &AppErrors{
		Code:            ErrInvalidBusinessRules,
		Map:             validate.FieldErrorsFlat(),
		Errors:          validate.Errors(),
		OriginFunc:      "CustomerProfile.Validate",
		FriendlyMessage: "Validation failed for the provided input.",
	}
}

func UnexpectedError(message string) *AppErrors {
	return &AppErrors{
		Code:            ErrInvalidBusinessRules,
		Errors:          []string{message},
		OriginFunc:      "UnexpectedError",
		FriendlyMessage: "Unexpected error.",
	}
}

func InvalidPasswordOrEmailError() *AppErrors {
	return &AppErrors{
		Code:            ErrInvalidBusinessRules,
		Errors:          []string{"invalid password or email"},
		OriginFunc:      "CustomerProfile.Validate",
		FriendlyMessage: "Validation failed for the provided input.",
	}
}
