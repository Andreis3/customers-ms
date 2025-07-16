package presentation

import (
	"github.com/andreis3/customers-ms/internal/app/services"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/infra/adapters/crypto"
	"github.com/andreis3/customers-ms/internal/infra/adapters/db"
	"github.com/andreis3/customers-ms/internal/infra/factories/app"
	"github.com/andreis3/customers-ms/internal/infra/factories/infra"
	"github.com/andreis3/customers-ms/internal/infra/repositories/repository"
	"github.com/andreis3/customers-ms/internal/presentation/http/handler"
	"github.com/andreis3/customers-ms/internal/presentation/http/routes"
)

func MakeCustomerRouter(
	postgres *db.Postgres,
	log adapter.Logger,
	prometheus adapter.Prometheus,
	tracer adapter.Tracer) *routes.CustomerRoutes {
	newCrypto := crypto.NewBcrypt()
	uowFactory := infra.NewUnitOfWorkFactory(postgres.Pool, prometheus, tracer)

	customerRepository := repository.NewCustomerRepository(postgres, prometheus, tracer)
	addressRepository := repository.NewAddressRepository(postgres, prometheus, tracer)
	customerService := services.NewCustomerService(customerRepository)
	command := app.NewCreateCustomerFactory(uowFactory, newCrypto, log, customerService, customerRepository, addressRepository, tracer)
	createCustomerHandler := handler.NewCreateCustomerHandler(command, prometheus, log, tracer)

	customerRoutes := routes.NewCustomer(createCustomerHandler, log, tracer)

	return customerRoutes
}
