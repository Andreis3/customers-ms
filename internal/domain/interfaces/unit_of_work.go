package interfaces

import apperror "github.com/andreis3/customers-ms/internal/domain/app-error"

type RepositoryFactory func(tx any) any

type UnitOfWork interface {
	Do(callback func(txUow UnitOfWork) *apperror.Error) *apperror.Error
	CustomerRepository() CustomerRepository
	AddressRepository() AddressRepository
}
