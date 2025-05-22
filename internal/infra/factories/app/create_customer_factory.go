package app

import (
	"github.com/andreis3/customers-ms/internal/app/commands"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/command"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/commons"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/service"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/uow"
)

type CreateCustomerFactory interface {
	Build() command.CreateCustomer
}

type createCustomerFactory struct {
	uow             uow.UnitOfWork
	crypt           adapter.Bcrypt
	log             commons.Logger
	customerService service.CustomerService
}

func NewCreateCustomerFactory(
	uow uow.UnitOfWork,
	crypto adapter.Bcrypt,
	log commons.Logger,
	customerService service.CustomerService,
) CreateCustomerFactory {
	return &createCustomerFactory{uow: uow, crypt: crypto, log: log, customerService: customerService}
}

func (f *createCustomerFactory) Build() command.CreateCustomer {
	return commands.NewCreateCustomer(f.uow, f.crypt, f.log, f.customerService)
}
