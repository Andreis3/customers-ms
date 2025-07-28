package router_factory

import (
	"github.com/andreis3/customers-ms/internal/app/services"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/infra/adapters/db"
	"github.com/andreis3/customers-ms/internal/infra/adapters/jwt"
	"github.com/andreis3/customers-ms/internal/infra/configs"
	"github.com/andreis3/customers-ms/internal/infra/factories/presentation/handler-factory"
	"github.com/andreis3/customers-ms/internal/infra/repositories/repository"
	"github.com/andreis3/customers-ms/internal/presentation/http/routes"
)

func MakeCustomerRouter(
	postgres *db.Postgres,
	redis *db.Redis,
	log adapter.Logger,
	prometheus adapter.Prometheus,
	tracer adapter.Tracer,
	conf *configs.Configs) *routes.CustomerRoutes {

	jwt := jwt.NewJWT(conf)
	customerRepository := repository.NewCustomerRepository(postgres, prometheus, tracer)
	authService := services.NewAuthService(jwt, customerRepository)
	getAddressHandler := handler_factory.NewGetCustomerAddresses(postgres, redis, log, prometheus, tracer, conf)

	createCustomerHandler := handler_factory.NewCreateCustomer(postgres, redis, log, prometheus, tracer, conf)
	customerRoutes := routes.NewCustomer(
		createCustomerHandler,
		getAddressHandler,
		authService,
		log,
		tracer,
	)
	return customerRoutes
}
