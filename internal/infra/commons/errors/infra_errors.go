package errors

import (
	domain_errors "github.com/andreis3/users-ms/internal/domain/errors"
)

const (
	ErrInternalProcessing domain_errors.ErrorCode = "IMF-500"
)

const (
	InternalServerError        = "internal server error"
	ServerErrorFriendlyMessage = "Server suffered an internal error"
)

// uow errors

func ErrorTransactionAlreadyExists() *domain_errors.AppErrors {
	return &domain_errors.AppErrors{
		Code:            ErrInternalProcessing,
		Errors:          []string{"Transaction already exists"},
		Cause:           InternalServerError,
		OriginFunc:      "UnitOfWork.Do",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorOpeningTransaction(err error) *domain_errors.AppErrors {
	return &domain_errors.AppErrors{
		Code:            ErrInternalProcessing,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "UnitOfWork.Do",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorRollBackTransactionEmpty() *domain_errors.AppErrors {
	return &domain_errors.AppErrors{
		Code:            ErrInternalProcessing,
		Errors:          []string{"Transaction is empty"},
		Cause:           InternalServerError,
		OriginFunc:      "UnitOfWork.Rollback",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorExecuteRollback(err error) *domain_errors.AppErrors {
	return &domain_errors.AppErrors{
		Code:            ErrInternalProcessing,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "UnitOfWork.Rollback",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

func ErrorCommitOrRollback(err error) *domain_errors.AppErrors {
	return &domain_errors.AppErrors{
		Code:            ErrInternalProcessing,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "UnitOfWork.CommitOrRollback",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}

// repository customers errors

func ErrorSaveCustomer(err error) *domain_errors.AppErrors {
	return &domain_errors.AppErrors{
		Code:            ErrInternalProcessing,
		Errors:          []string{err.Error()},
		Cause:           InternalServerError,
		OriginFunc:      "CustomerRepository.SaveCustomer",
		FriendlyMessage: ServerErrorFriendlyMessage,
	}
}
