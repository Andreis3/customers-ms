package interfaces

import "github.com/andreis3/users-ms/internal/domain/errors"

type RepositoryFactory func(tx any) any

type UnitOfWork interface {
	Register(name string, callback RepositoryFactory)
	GetRepository(name string) any
	Do(callback func(uow UnitOfWork) *errors.AppErrors) *errors.AppErrors
	CommitOrRollback() *errors.AppErrors
	Rollback() *errors.AppErrors
}
