package factories

import (
	"github.com/andreis3/users-ms/internal/domain/errors"
	"github.com/andreis3/users-ms/internal/domain/interfaces"
	"github.com/andreis3/users-ms/internal/infra/uow"
)

type CreateCustomer struct {
	CustomerRepo interfaces.CustomerRepository
	AddressRepo  interfaces.AddressRepository
}

func LoadCustomerFactory(unitOfWork interfaces.UnitOfWork) (*CreateCustomer, *errors.AppErrors) {
	customerRepo, ok := unitOfWork.GetRepository(uow.CustomerRepository).(interfaces.CustomerRepository)
	if !ok {
		return nil, errors.UnexpectedError("unexpected repository")
	}
	addressRepo, ok := unitOfWork.GetRepository(uow.AddressRepository).(interfaces.AddressRepository)
	if !ok {
		return nil, errors.UnexpectedError("unexpected repository")
	}
	return &CreateCustomer{
		CustomerRepo: customerRepo,
		AddressRepo:  addressRepo,
	}, nil
}
