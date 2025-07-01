package errors

import (
	"errors"
	"strings"

	"github.com/andreis3/customers-ms/internal/domain/validator"
)

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

func (e *Error) Error() string {
	return string(e.Code) + ": " + e.FriendlyMessage
}

func (e *Error) Unwrap() error {
	if len(e.Errors) > 0 {
		return errors.New(strings.Join(e.Errors, "; "))
	}
	return nil
}

func As(err error, target any) bool {
	return errors.As(err, target)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func (e *Error) Is(target error) bool {
	var t *Error
	ok := errors.As(target, &t)
	if !ok {
		return false
	}
	return e.Code == t.Code
}

func New(code Code, originFunc, friendlyMessage string, errs ...string) *Error {
	return &Error{
		Code:            code,
		Errors:          errs,
		OriginFunc:      originFunc,
		FriendlyMessage: friendlyMessage,
	}
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
	return New(BadRequestCode, "UnexpectedError", "Unexpected error.", message)
}

func InvalidPasswordOrEmailError() *Error {
	return New(BadRequestCode, "CustomerProfile.Validate", "Validation failed for the provided input.", "invalid password or email")
}

func ErrCustomerAlreadyExists() *Error {
	return New(BadRequestCode, "CustomerProfile.Validate", "Already exists a customer with this email.", "customer already exists")
}

/********UnitOfWork Errors********/
func ErrorTransactionAlreadyExists() *Error {
	return New(InternalServerErrorCode, "UnitOfWork.Do", ServerErrorFriendlyMessage, "Transaction already exists")
}

func ErrorOpeningTransaction(err error) *Error {
	return New(InternalServerErrorCode, "UnitOfWork.Do", ServerErrorFriendlyMessage, err.Error())
}

func ErrorRollBackTransactionEmpty() *Error {
	return New(InternalServerErrorCode, "UnitOfWork.Rollback", ServerErrorFriendlyMessage, "Transaction is empty")
}

func ErrorExecuteRollback(err error) *Error {
	return New(InternalServerErrorCode, "UnitOfWork.Rollback", ServerErrorFriendlyMessage, err.Error())
}

func ErrorCommitOrRollback(err error) *Error {
	return New(InternalServerErrorCode, "UnitOfWork.CommitOrRollback", ServerErrorFriendlyMessage, err.Error())
}

/********Repository Errors********/
func ErrorSaveCustomer(err error) *Error {
	return New(InternalServerErrorCode, "CustomerRepository.SaveCustomer", ServerErrorFriendlyMessage, err.Error())
}

func ErrorFindCustomerByEmail(err error) *Error {
	return New(InternalServerErrorCode, "CustomerRepository.FindCustomerByEmail", ServerErrorFriendlyMessage, err.Error())
}

func ErrorCreatedBatchAddress(err error) *Error {
	return New(InternalServerErrorCode, "CustomerRepository.CreatedBatchAddress", ServerErrorFriendlyMessage, err.Error())
}

/********Bcrypt Errors********/
func ErrorHashPassword(err error) *Error {
	return New(InternalServerErrorCode, "Bcrypt.Hash", ServerErrorFriendlyMessage, err.Error())
}

func ErrorCompareHash(err error) *Error {
	return New(InternalServerErrorCode, "Bcrypt.CompareHash", ServerErrorFriendlyMessage, err.Error())
}

/********JWT Errors********/
func ErrorCreateToken(err error) *Error {
	return New(InternalServerErrorCode, "JWT.CreateToken", ServerErrorFriendlyMessage, err.Error())
}

func ErrorValidateToken(err error) *Error {
	return New(InternalServerErrorCode, "JWT.ValidateToken", ServerErrorFriendlyMessage, err.Error())
}

func ErrorInvalidTokenAlgorithmError() *Error {
	return New(InternalServerErrorCode, "JWT.ValidateToken", ServerErrorFriendlyMessage, "invalid token algorithm")
}

func ErrorRefreshToken(err error) *Error {
	return New(InternalServerErrorCode, "JWT.RefreshToken", ServerErrorFriendlyMessage, err.Error())
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
