package presentation

import (
	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/domain/services"
	"github.com/andreis3/customers-ms/internal/infra/adapters/crypto"
	"github.com/andreis3/customers-ms/internal/infra/adapters/jwt"
	"github.com/andreis3/customers-ms/internal/infra/configs"
	"github.com/andreis3/customers-ms/internal/infra/factories/app"
	"github.com/andreis3/customers-ms/internal/infra/repositories/repository"
	"github.com/andreis3/customers-ms/internal/presentation/http/handler"
	"github.com/andreis3/customers-ms/internal/presentation/http/routes"
)

func MakeAuthRouter(postgres adapter.Postgres, log adapter.Logger, prometheus adapter.Prometheus, conf *configs.Configs, tracer adapter.Tracer) *routes.LoginRoutes {
	customerRepository := repository.NewCustomerRepository(postgres, prometheus, tracer)
	tokenService := jwt.NewJWT(conf)
	authService := services.NewAuthService(tokenService)
	bcrypt := crypto.NewBcrypt()
	commands := app.NewAuthenticateCustomerFactory(log, customerRepository, authService, bcrypt, tracer)
	authHandler := handler.NewGenerateTokenHandler(log, prometheus, tracer, commands)
	return routes.NewLoginRoutes(log, authHandler, tracer)
}
