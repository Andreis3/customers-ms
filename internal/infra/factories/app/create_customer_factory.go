package app

import (
	"github.com/andreis3/customers-ms/internal/app/commands"
	"github.com/andreis3/customers-ms/internal/app/services"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/command"
	"github.com/andreis3/customers-ms/internal/infra/adapters/db"
	"github.com/andreis3/customers-ms/internal/infra/factories/infra"
	"github.com/andreis3/customers-ms/internal/infra/repositories/repository"
)

func NewCreateCustomerFactory(
	db *db.Postgres,
	crypto adapter.Bcrypt,
	log adapter.Logger,
	tracer adapter.Tracer,
	metrics adapter.Prometheus,
) command.CreateCustomer {
	customerRepository := repository.NewCustomerRepository(db, metrics, tracer)
	addressRepository := repository.NewAddressRepository(db, metrics, tracer)
	uowFactory := infra.NewUnitOfWorkFactory(db.Pool, metrics, tracer)
	customerService := services.NewCustomerService(customerRepository)
	return commands.NewCreateCustomer(uowFactory, crypto, log, customerService, customerRepository, addressRepository, tracer)
}
