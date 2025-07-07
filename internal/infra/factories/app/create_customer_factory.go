package app

import (
	"context"

	"github.com/andreis3/customers-ms/internal/app/commands"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/command"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/commons"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/postgres"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/service"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/uow"
)

type createCustomerFactory struct {
	uowFactory         func(ctx context.Context) uow.UnitOfWork
	crypt              adapter.Bcrypt
	log                commons.Logger
	customerService    service.CustomerService
	customerRepository postgres.CustomerRepository
	addressRepository  postgres.AddressRepository
	tracer             adapter.Tracer
}

func NewCreateCustomerFactory(
	uowFactory func(ctx context.Context) uow.UnitOfWork,
	crypto adapter.Bcrypt,
	log commons.Logger,
	customerService service.CustomerService,
	customerRepository postgres.CustomerRepository,
	addressRepository postgres.AddressRepository,
	tracer adapter.Tracer,
) command.CreateCustomer {
	return commands.NewCreateCustomer(uowFactory, crypto, log, customerService, customerRepository, addressRepository, tracer)
}
