package uow

import (
	"context"

	apperror "github.com/andreis3/customers-ms/internal/domain/app-error"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/repository"
)

type RepositoryFactory func(tx any) any

type UnitOfWork interface {
	Do(ctx context.Context, callback func(txUow UnitOfWork) *apperror.Error) *apperror.Error
	CustomerRepository() repository.CustomerRepository
	AddressRepository() repository.AddressRepository
}
