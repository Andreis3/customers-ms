package presentation

import (
	"github.com/andreis3/customers-ms/internal/app/services"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/infra/adapters/crypto"
	"github.com/andreis3/customers-ms/internal/infra/adapters/db"
	"github.com/andreis3/customers-ms/internal/infra/adapters/jwt"
	"github.com/andreis3/customers-ms/internal/infra/configs"
	"github.com/andreis3/customers-ms/internal/infra/factories/app"
	"github.com/andreis3/customers-ms/internal/infra/repositories/repository"
	"github.com/andreis3/customers-ms/internal/presentation/http/handler"
	"github.com/andreis3/customers-ms/internal/presentation/http/routes"
)

func MakeCustomerRouter(
	postgres *db.Postgres,
	log adapter.Logger,
	prometheus adapter.Prometheus,
	tracer adapter.Tracer,
	conf *configs.Configs) *routes.CustomerRoutes {
	newCrypto := crypto.NewBcrypt()
	command := app.NewCreateCustomerFactory(
		postgres,
		newCrypto,
		log,
		tracer,
		prometheus,
	)
	jwt := jwt.NewJWT(conf)
	customerRepository := repository.NewCustomerRepository(postgres, prometheus, tracer)
	authService := services.NewAuthService(jwt, customerRepository)
	query := app.NewGetCustomerAddressesFactory(postgres, log, tracer, prometheus)
	getAddressHandler := handler.NewGetAddressHandler(query, log, prometheus, tracer)

	createCustomerHandler := handler.NewCreateCustomerHandler(command, prometheus, log, tracer)
	customerRoutes := routes.NewCustomer(
		createCustomerHandler,
		*getAddressHandler,
		authService,
		log,
		tracer,
	)
	return customerRoutes
}
