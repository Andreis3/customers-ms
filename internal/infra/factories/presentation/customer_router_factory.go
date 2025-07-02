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

func MakeCustomerRouter(connPostgres *postegres.Postgres, log commons.Logger, prometheus adapter.Prometheus) *routes.CustomerRoutes {
	pool := connPostgres.Pool
	newCrypto := crypto.NewBcrypt()
	uowFactory := app.NewUnitOfWorkFactory(pool, prometheus)

	customerRepository := repository.NewCustomerRepository(pool, prometheus)
	addressRepository := repository.NewAddressRepository(pool, prometheus)
	customerService := services.NewCustomerService(customerRepository)
	command := app.NewCreateCustomerFactory(uowFactory, newCrypto, log, customerService, customerRepository, addressRepository)
	createCustomerHandler := customer.NewCreateCustomerHandler(command, prometheus, log)

	customerRoutes := routes.NewCustomer(createCustomerHandler, log)

	return customerRoutes
}
