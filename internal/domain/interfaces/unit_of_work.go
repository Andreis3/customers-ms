package interfaces

import (
	"github.com/andreis3/users-ms/internal/domain/apperrors"
)

type RepositoryFactory func(tx any) any

type UnitOfWork interface {
	Do(callback func(uow UnitOfWork) *apperrors.AppErrors) *apperrors.AppErrors
	CommitOrRollback() *apperrors.AppErrors
	Rollback() *apperrors.AppErrors
	TXExecute() InstructionPostgres
}
