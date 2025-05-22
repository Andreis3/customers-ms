package presentation

import (
	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/commons"
	"github.com/andreis3/customers-ms/internal/domain/services"
	"github.com/andreis3/customers-ms/internal/infra/adapters/crypto"
	"github.com/andreis3/customers-ms/internal/infra/adapters/db/postegres"
	"github.com/andreis3/customers-ms/internal/infra/adapters/jwt"
	"github.com/andreis3/customers-ms/internal/infra/configs"

	"github.com/andreis3/customers-ms/internal/infra/repositories/postgres/repository"
	"github.com/andreis3/customers-ms/internal/presentation/http/handler/auth"
	"github.com/andreis3/customers-ms/internal/presentation/http/routes"
)

func MakeAuthRouter(connPostgres *postegres.Postgres, log commons.Logger, prometheus adapter.Prometheus, conf *configs.Configs) *routes.AuthRoutes {
	pool := connPostgres.Pool
	customerRepository := repository.NewCustomerRepository(pool, prometheus)
	tokenService := jwt.NewJWT(conf)
	authService := services.NewAuthService(tokenService)
	bcrypt := crypto.NewBcrypt()
	authHandler := auth.NewGenerateTokenHandler(log, prometheus, customerRepository, authService, bcrypt)
	return routes.NewAuthRoutes(log, authHandler)
}
