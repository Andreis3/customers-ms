package errors

import "github.com/andreis3/users-ms/internal/domain/validator"

type ErrorCode string

const (
	ErrInvalidBusinessRules ErrorCode = "DM-400"
	ErrResourceNotFound     ErrorCode = "DM-404"
)

func InvalidCustomerError(validate *validator.Validator) *AppErrors {
	return &AppErrors{
		Code:            ErrInvalidBusinessRules,
		Errors:          validate.Errors(),
		OriginFunc:      "CustomerProfile.Validate",
		FriendlyMessage: "Validation failed for the provided input.",
	}
}
