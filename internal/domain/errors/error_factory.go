package errors

import "github.com/andreis3/customers-ms/internal/domain/validator"

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
	input := InputError{
		Code:            InternalServerError,
		Errors:          []string{message},
		OriginFunc:      "UnexpectedError",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
	return New(input)
}

func InvalidPasswordOrEmailError() *Error {
	input := InputError{
		Code:            BadRequestCode,
		Errors:          []string{"Validation failed for the provided input."},
		OriginFunc:      "CustomerProfile.Validate",
		FriendlyMessage: "invalid password or email",
	}
	return New(input)
}

func ErrCustomerAlreadyExists() *Error {
	input := InputError{
		Code:            BadRequestCode,
		Errors:          []string{"Already exists a customer with this email."},
		OriginFunc:      "CustomerProfile.Validate",
		FriendlyMessage: "Customer already exists",
	}
	return New(input)
}

/********UnitOfWork Errors********/
func ErrorTransactionAlreadyExists() *Error {
	input := InputError{
		Code:            InternalServerErrorCode,
		Errors:          []string{ServerErrorFriendlyMessage},
		OriginFunc:      "UnitOfWork.Do",
		FriendlyMessage: "Transaction already exists",
	}
	return New(input)
}

func ErrorOpeningTransaction(err error) *Error {
	input := InputError{
		Code:            InternalServerErrorCode,
		Errors:          []string{err.Error()},
		OriginFunc:      "UnitOfWork.Do",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
	return New(input)
}

func ErrorRollBackTransactionEmpty() *Error {
	input := InputError{
		Code:            InternalServerErrorCode,
		Errors:          []string{"Transaction is empty"},
		OriginFunc:      "UnitOfWork.Rollback",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
	return New(input)
}

func ErrorExecuteRollback(err error) *Error {
	input := InputError{
		Code:            InternalServerErrorCode,
		Errors:          []string{err.Error()},
		OriginFunc:      "UnitOfWork.Rollback",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
	return New(input)
}

func ErrorCommitOrRollback(err error) *Error {
	input := InputError{
		Code:            InternalServerErrorCode,
		Errors:          []string{err.Error()},
		OriginFunc:      "UnitOfWork.CommitOrRollback",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
	return New(input)
}

/********Repository Errors********/
func ErrorSaveCustomer(err error) *Error {
	input := InputError{
		Code:            InternalServerErrorCode,
		Errors:          []string{err.Error()},
		OriginFunc:      "CustomerRepository.SaveCustomer",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
	return New(input)
}

func ErrorFindCustomerByEmail(err error) *Error {
	input := InputError{
		Code:            InternalServerErrorCode,
		Errors:          []string{err.Error()},
		OriginFunc:      "CustomerRepository.FindCustomerByEmail",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
	return New(input)
}

func ErrorCreatedBatchAddress(err error) *Error {
	input := InputError{
		Code:            InternalServerErrorCode,
		Errors:          []string{err.Error()},
		OriginFunc:      "CustomerRepository.CreatedBatchAddress",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
	return New(input)
}

/********Bcrypt Errors********/
func ErrorHashPassword(err error) *Error {
	input := InputError{
		Code:            InternalServerErrorCode,
		Errors:          []string{err.Error()},
		OriginFunc:      "Bcrypt.Hash",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
	return New(input)
}

func ErrorCompareHash(err error) *Error {
	input := InputError{
		Code:            InternalServerErrorCode,
		Errors:          []string{err.Error()},
		OriginFunc:      "Bcrypt.CompareHash",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
	return New(input)
}

/********JWT Errors********/
func ErrorCreateToken(err error) *Error {
	input := InputError{
		Code:            InternalServerErrorCode,
		Errors:          []string{err.Error()},
		OriginFunc:      "JWT.CreateToken",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
	return New(input)
}

func ErrorValidateToken(err error) *Error {
	input := InputError{
		Code:            InternalServerErrorCode,
		Errors:          []string{err.Error()},
		OriginFunc:      "JWT.ValidateToken",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
	return New(input)
}

func ErrorInvalidTokenAlgorithmError() *Error {
	input := InputError{
		Code:            InternalServerErrorCode,
		Errors:          []string{"invalid token algorithm"},
		OriginFunc:      "JWT.ValidateToken",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
	return New(input)
}

func ErrorRefreshToken(err error) *Error {
	input := InputError{
		Code:            InternalServerErrorCode,
		Errors:          []string{err.Error()},
		OriginFunc:      "JWT.RefreshToken",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
	return New(input)
}

func ErrorValidateTokenMessage(message string) *Error {
	input := InputError{
		Code:            InternalServerErrorCode,
		Errors:          []string{message},
		OriginFunc:      "JWT.ValidateToken",
		FriendlyMessage: InvalidCredentialsMessage,
	}
	return New(input)
}

func ErrorInvalidCredentials() *Error {
	input := InputError{
		Code:            BadRequestCode,
		Errors:          []string{"invalid credentials"},
		OriginFunc:      "JWT.ValidateToken",
		FriendlyMessage: InvalidCredentialsMessage,
	}
	return New(input)
}

/********Decoder Errors********/
func ErrorJSONSyntaxError(err error) *Error {
	input := InputError{
		Code:            BadRequestCode,
		Errors:          []string{err.Error()},
		Cause:           "json syntax error",
		OriginFunc:      "json.Unmarshal",
		FriendlyMessage: "json syntax error",
	}
	return New(input)
}

func ErrorJSONUnmarshalTypeError(err error) *Error {
	input := InputError{
		Code:            BadRequestCode,
		Errors:          []string{err.Error()},
		Cause:           "json unmarshal type error",
		OriginFunc:      "json.Unmarshal",
		FriendlyMessage: "json unmarshal type error",
	}
	return New(input)
}

func ErrorJSON(err error) *Error {
	input := InputError{
		Code:            BadRequestCode,
		Errors:          []string{err.Error()},
		Cause:           "json error",
		OriginFunc:      "json.Unmarshal",
		FriendlyMessage: "json error",
	}
	return New(input)
}
