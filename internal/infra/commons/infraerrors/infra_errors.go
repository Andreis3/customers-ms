package infraerrors

import (
	"github.com/andreis3/customers-ms/internal/domain/apperrors"
)

const (
	ErrInternalProcessing apperrors.ErrorCode = "IMF-500"
)

const (
	InternalServerError        = "internal server error"
	ServerErrorFriendlyMessage = "Server suffered an internal error"
)

// uow errors

func ErrorTransactionAlreadyExists() *apperrors.AppErrors {
	return &apperrors.AppErrors{
		Code:            ErrInternalProcessing,
		Errors:          []string{"Transaction already exists"},
		Cause:           InternalServerError,
		OriginFunc:      "UnitOfWork.Do",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorOpeningTransaction(err error) *apperrors.AppErrors {
	return &apperrors.AppErrors{
		Code:            ErrInternalProcessing,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "UnitOfWork.Do",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorRollBackTransactionEmpty() *apperrors.AppErrors {
	return &apperrors.AppErrors{
		Code:            ErrInternalProcessing,
		Errors:          []string{"Transaction is empty"},
		Cause:           InternalServerError,
		OriginFunc:      "UnitOfWork.Rollback",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorExecuteRollback(err error) *apperrors.AppErrors {
	return &apperrors.AppErrors{
		Code:            ErrInternalProcessing,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "UnitOfWork.Rollback",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorCommitOrRollback(err error) *apperrors.AppErrors {
	return &apperrors.AppErrors{
		Code:            ErrInternalProcessing,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "UnitOfWork.CommitOrRollback",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

// repository customers errors

func ErrorSaveCustomer(err error) *apperrors.AppErrors {
	return &apperrors.AppErrors{
		Code:            ErrInternalProcessing,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "CustomerRepository.SaveCustomer",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorCreatedBatchAddress(err error) *apperrors.AppErrors {
	return &apperrors.AppErrors{
		Code:            ErrInternalProcessing,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "CustomerRepository.CreatedBatchAddress",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

// bcrypt
func ErrorHashPassword(err error) *apperrors.AppErrors {
	return &apperrors.AppErrors{
		Code:            ErrInternalProcessing,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "Bcrypt.Hash",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorCompareHash(err error) *apperrors.AppErrors {
	return &apperrors.AppErrors{
		Code:            ErrInternalProcessing,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "Bcrypt.CompareHash",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}
