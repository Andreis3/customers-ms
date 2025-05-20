package apperror

import "github.com/andreis3/customers-ms/internal/domain/validator"

type Code string

const (
	BadRequestCode          Code = "DM-400"
	NotFoundCode            Code = "DM-404"
	InternalServerErrorCode Code = "IMF-500"
	UnauthorizedCode        Code = "DM-401"
	ForbiddenCode           Code = "DM-403"
	ConflictCode            Code = "DM-409"
	UnprocessableEntityCode Code = "DM-422"
)

const (
	InternalServerError        = "Internal server error"
	ServerErrorFriendlyMessage = "Internal server error"
	InvalidCredentialsMessage  = "Invalid credentials"
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
		Code:            BadRequestCode,
		Map:             validate.FieldErrorsGrouped(),
		Errors:          validate.Errors(),
		OriginFunc:      "CustomerProfile.Validate",
		FriendlyMessage: "Validation failed for the provided input.",
	}
}

func UnexpectedError(message string) *Error {
	return &Error{
		Code:            BadRequestCode,
		Errors:          []string{message},
		OriginFunc:      "UnexpectedError",
		FriendlyMessage: "Unexpected error.",
	}
}

func InvalidPasswordOrEmailError() *Error {
	return &Error{
		Code:            BadRequestCode,
		Errors:          []string{"invalid password or email"},
		OriginFunc:      "CustomerProfile.Validate",
		FriendlyMessage: "Validation failed for the provided input.",
	}
}

func ErrCustomerAlreadyExists() *Error {
	return &Error{
		Code:            BadRequestCode,
		Errors:          []string{"customer already exists"},
		OriginFunc:      "CustomerProfile.Validate",
		FriendlyMessage: "Already exists a customer with this email.",
	}
}

/********UnitOfWork Errors********/
func ErrorTransactionAlreadyExists() *Error {
	return &Error{
		Code:            InternalServerErrorCode,
		Errors:          []string{"Transaction already exists"},
		Cause:           InternalServerError,
		OriginFunc:      "UnitOfWork.Do",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorOpeningTransaction(err error) *Error {
	return &Error{
		Code:            InternalServerErrorCode,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "UnitOfWork.Do",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorRollBackTransactionEmpty() *Error {
	return &Error{
		Code:            InternalServerErrorCode,
		Errors:          []string{"Transaction is empty"},
		Cause:           InternalServerError,
		OriginFunc:      "UnitOfWork.Rollback",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorExecuteRollback(err error) *Error {
	return &Error{
		Code:            InternalServerErrorCode,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "UnitOfWork.Rollback",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorCommitOrRollback(err error) *Error {
	return &Error{
		Code:            InternalServerErrorCode,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "UnitOfWork.CommitOrRollback",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

/********Repository Errors********/
func ErrorSaveCustomer(err error) *Error {
	return &Error{
		Code:            InternalServerErrorCode,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "CustomerRepository.SaveCustomer",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorFindCustomerByEmail(err error) *Error {
	return &Error{
		Code:            InternalServerErrorCode,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "CustomerRepository.FindCustomerByEmail",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorCreatedBatchAddress(err error) *Error {
	return &Error{
		Code:            InternalServerErrorCode,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "CustomerRepository.CreatedBatchAddress",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

/********Bcrypt Errors********/
func ErrorHashPassword(err error) *Error {
	return &Error{
		Code:            InternalServerErrorCode,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "Bcrypt.Hash",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorCompareHash(err error) *Error {
	return &Error{
		Code:            InternalServerErrorCode,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "Bcrypt.CompareHash",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

/********JWT Errors********/
func ErrorCreateToken(err error) *Error {
	return &Error{
		Code:            InternalServerErrorCode,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "JWT.CreateToken",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorValidateToken(err error) *Error {
	return &Error{
		Code:            InternalServerErrorCode,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "JWT.ValidateToken",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorInvalidTokenAlgorithmError() *Error {
	return &Error{
		Code:            InternalServerErrorCode,
		Errors:          []string{"invalid token algorithm"},
		Cause:           InternalServerError,
		OriginFunc:      "JWT.ValidateToken",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorRefreshToken(err error) *Error {
	return &Error{
		Code:            InternalServerErrorCode,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "JWT.RefreshToken",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorValidateTokenMessage(message string) *Error {
	return &Error{
		Code:            InternalServerErrorCode,
		Errors:          []string{message},
		Cause:           InvalidCredentialsMessage,
		OriginFunc:      "JWT.ValidateToken",
		FriendlyMessage: InvalidCredentialsMessage,
	}
}

func ErrorInvalidCredentials() *Error {
	return &Error{
		Code:            UnauthorizedCode,
		Errors:          []string{"invalid credentials"},
		Cause:           InvalidCredentialsMessage,
		OriginFunc:      "JWT.ValidateToken",
		FriendlyMessage: InvalidCredentialsMessage,
	}
}
