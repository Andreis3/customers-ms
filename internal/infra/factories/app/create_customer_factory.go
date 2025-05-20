package app

import (
	"github.com/andreis3/customers-ms/internal/app/commands"
	"github.com/andreis3/customers-ms/internal/domain/interfaces"
)

type createCustomerFactory struct {
	uow             interfaces.UnitOfWork
	crypt           interfaces.Bcrypt
	log             interfaces.Logger
	customerService interfaces.CustomerService
}

type ICreateCustomerFactory interface {
	Build() commands.ICreateCustomer
}

func NewCreateCustomerFactory(
	uow interfaces.UnitOfWork,
	crypto interfaces.Bcrypt,
	log interfaces.Logger,
	customerService interfaces.CustomerService,
) ICreateCustomerFactory {
	return &createCustomerFactory{uow: uow, crypt: crypto, log: log, customerService: customerService}
}

func (f *createCustomerFactory) Build() commands.ICreateCustomer {
	return commands.NewCreateCustomer(f.uow, f.crypt, f.log, f.customerService)
}
