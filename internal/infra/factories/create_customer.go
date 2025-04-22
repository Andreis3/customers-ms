package factories

import (
	"github.com/andreis3/users-ms/internal/domain/apperrors"
	"github.com/andreis3/users-ms/internal/domain/interfaces"
	"github.com/andreis3/users-ms/internal/infra/uow"
)

type CreateCustomer struct {
	Customer interfaces.CustomerRepository
	Address  interfaces.AddressRepository
}

func LoadCustomerFactory(unitOfWork interfaces.UnitOfWork) (*CreateCustomer, *apperrors.AppErrors) {
	customerRepo, ok := unitOfWork.GetRepository(uow.CustomerRepository).(interfaces.CustomerRepository)
	if !ok {
		return nil, apperrors.UnexpectedError("unexpected repository")
	}
	addressRepo, ok := unitOfWork.GetRepository(uow.AddressRepository).(interfaces.AddressRepository)
	if !ok {
		return nil, apperrors.UnexpectedError("unexpected repository")
	}
	return &CreateCustomer{
		Customer: customerRepo,
		Address:  addressRepo,
	}, nil
}
