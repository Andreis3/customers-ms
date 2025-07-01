package uow

import (
	"context"

	"github.com/andreis3/customers-ms/internal/domain/errors"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/postgres"
)

type RepositoryFactory func(tx any) any

type UnitOfWork interface {
	Do(ctx context.Context, callback func(txUow UnitOfWork) *errors.Error) *errors.Error
	CustomerRepository() postgres.CustomerRepository
	AddressRepository() postgres.AddressRepository
}
