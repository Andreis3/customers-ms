package apperror

import "github.com/andreis3/customers-ms/internal/domain/validator"

type Code string

const (
	ErrInvalidBusinessRules Code = "DM-400"
	ErrResourceNotFound     Code = "DM-404"
	ErrInternalProcessing   Code = "IMF-500"
	ErrUnauthorized         Code = "DM-401"
	ErrForbidden            Code = "DM-403"
)

const (
	InternalServerError        = "Internal server error"
	ServerErrorFriendlyMessage = "Internal server error"
)

type Error struct {
	Code            Code
	Errors          []string
	Map             map[string]any
	OriginFunc      string
	Cause           string
	FriendlyMessage string
}

func (e Error) Error() string {
	return string(e.Code) + ": " + e.FriendlyMessage
}

/********Domain Errors********/

func InvalidCustomerError(validate *validator.Validator) *Error {
	return &Error{
		Code:            ErrInvalidBusinessRules,
		Map:             validate.FieldErrorsGrouped(),
		Errors:          validate.Errors(),
		OriginFunc:      "CustomerProfile.Validate",
		FriendlyMessage: "Validation failed for the provided input.",
	}
}

func UnexpectedError(message string) *Error {
	return &Error{
		Code:            ErrInvalidBusinessRules,
		Errors:          []string{message},
		OriginFunc:      "UnexpectedError",
		FriendlyMessage: "Unexpected error.",
	}
}

func InvalidPasswordOrEmailError() *Error {
	return &Error{
		Code:            ErrInvalidBusinessRules,
		Errors:          []string{"invalid password or email"},
		OriginFunc:      "CustomerProfile.Validate",
		FriendlyMessage: "Validation failed for the provided input.",
	}
}

func ErrCustomerAlreadyExists() *Error {
	return &Error{
		Code:            ErrInvalidBusinessRules,
		Errors:          []string{"customer already exists"},
		OriginFunc:      "CustomerProfile.Validate",
		FriendlyMessage: "Already exists a customer with this email.",
	}
}

/********UnitOfWork Errors********/
func ErrorTransactionAlreadyExists() *Error {
	return &Error{
		Code:            ErrInternalProcessing,
		Errors:          []string{"Transaction already exists"},
		Cause:           InternalServerError,
		OriginFunc:      "UnitOfWork.Do",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorOpeningTransaction(err error) *Error {
	return &Error{
		Code:            ErrInternalProcessing,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "UnitOfWork.Do",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorRollBackTransactionEmpty() *Error {
	return &Error{
		Code:            ErrInternalProcessing,
		Errors:          []string{"Transaction is empty"},
		Cause:           InternalServerError,
		OriginFunc:      "UnitOfWork.Rollback",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorExecuteRollback(err error) *Error {
	return &Error{
		Code:            ErrInternalProcessing,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "UnitOfWork.Rollback",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorCommitOrRollback(err error) *Error {
	return &Error{
		Code:            ErrInternalProcessing,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "UnitOfWork.CommitOrRollback",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

/********Repository Errors********/
func ErrorSaveCustomer(err error) *Error {
	return &Error{
		Code:            ErrInternalProcessing,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "CustomerRepository.SaveCustomer",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorFindCustomerByEmail(err error) *Error {
	return &Error{
		Code:            ErrInternalProcessing,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "CustomerRepository.FindCustomerByEmail",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorCreatedBatchAddress(err error) *Error {
	return &Error{
		Code:            ErrInternalProcessing,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "CustomerRepository.CreatedBatchAddress",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

/********Bcrypt Errors********/
func ErrorHashPassword(err error) *Error {
	return &Error{
		Code:            ErrInternalProcessing,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "Bcrypt.Hash",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorCompareHash(err error) *Error {
	return &Error{
		Code:            ErrInternalProcessing,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "Bcrypt.CompareHash",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

/********JWT Errors********/
func ErrorCreateToken(err error) *Error {
	return &Error{
		Code:            ErrInternalProcessing,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "JWT.CreateToken",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorValidateToken(err error) *Error {
	return &Error{
		Code:            ErrInternalProcessing,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "JWT.ValidateToken",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorInvalidTokenAlgorithmError() *Error {
	return &Error{
		Code:            ErrInternalProcessing,
		Errors:          []string{"invalid token algorithm"},
		Cause:           InternalServerError,
		OriginFunc:      "JWT.ValidateToken",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorRefreshToken(err error) *Error {
	return &Error{
		Code:            ErrInternalProcessing,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "JWT.RefreshToken",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorValidateTokenMessage(message string) *Error {
	return &Error{
		Code:            ErrInternalProcessing,
		Errors:          []string{message},
		Cause:           InternalServerError,
		OriginFunc:      "JWT.ValidateToken",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}
