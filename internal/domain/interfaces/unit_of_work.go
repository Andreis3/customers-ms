package interfaces

import (
	"github.com/andreis3/customers-ms/internal/domain/apperrors"
)

type RepositoryFactory func(tx any) any

type UnitOfWork interface {
	Do(callback func(txUow UnitOfWork) *apperrors.AppErrors) *apperrors.AppErrors
	CustomerRepository() CustomerRepository
	AddressRepository() AddressRepository
}
