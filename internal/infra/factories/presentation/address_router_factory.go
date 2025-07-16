package presentation

import (
	"github.com/andreis3/customers-ms/internal/app/services"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/infra/adapters/db"
	"github.com/andreis3/customers-ms/internal/infra/adapters/jwt"
	"github.com/andreis3/customers-ms/internal/infra/configs"
	"github.com/andreis3/customers-ms/internal/infra/repositories/repository"
	"github.com/andreis3/customers-ms/internal/presentation/http/handler"
	"github.com/andreis3/customers-ms/internal/presentation/http/routes"
)

func MakeAddressRouter(
	postgres *db.Postgres,
	log adapter.Logger,
	prometheus adapter.Prometheus,
	tracer adapter.Tracer,
	conf *configs.Configs) *routes.AddressRoutes {

	customerRepository := repository.NewCustomerRepository(postgres, prometheus, tracer)
	jwt := jwt.NewJWT(conf)
	authService := services.NewAuthService(jwt, customerRepository)
	getAddressHandler := handler.NewGetAddressHandler()

	customerRoutes := routes.NewAddressRoutes(getAddressHandler, authService, log, tracer)

	return customerRoutes
}
