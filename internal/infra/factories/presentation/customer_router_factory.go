package presentation

import (
	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/commons"
	"github.com/andreis3/customers-ms/internal/domain/services"
	"github.com/andreis3/customers-ms/internal/infra/adapters/crypto"
	"github.com/andreis3/customers-ms/internal/infra/adapters/db/postegres"
	"github.com/andreis3/customers-ms/internal/infra/factories/app"
	"github.com/andreis3/customers-ms/internal/infra/repositories/postgres/repository"
	"github.com/andreis3/customers-ms/internal/presentation/http/handler/customer"
	"github.com/andreis3/customers-ms/internal/presentation/http/routes"
)

func MakeCustomerRouter(connPostgres *postegres.Postgres, log commons.Logger, prometheus adapter.Prometheus, tracer adapter.Tracer) *routes.CustomerRoutes {
	pool := connPostgres.Pool
	newCrypto := crypto.NewBcrypt()
	uowFactory := app.NewUnitOfWorkFactory(pool, prometheus, tracer)

	customerRepository := repository.NewCustomerRepository(pool, prometheus, tracer)
	addressRepository := repository.NewAddressRepository(pool, prometheus, tracer)
	customerService := services.NewCustomerService(customerRepository)
	command := app.NewCreateCustomerFactory(uowFactory, newCrypto, log, customerService, customerRepository, addressRepository, tracer)
	createCustomerHandler := customer.NewCreateCustomerHandler(command, prometheus, log, tracer)

	customerRoutes := routes.NewCustomer(createCustomerHandler, log, tracer)

	return customerRoutes
}
