package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/commons"
	"github.com/andreis3/customers-ms/internal/infra/adapters/db/postegres"
	"github.com/andreis3/customers-ms/internal/infra/configs"
	"github.com/andreis3/customers-ms/internal/infra/factories/presentation"
	"github.com/andreis3/customers-ms/internal/presentation/http/routes"
)

type SetupRoutesInput struct {
	Mux          *chi.Mux
	ConnPostgres *postegres.Postgres
	Log          commons.Logger
	Prometheus   adapter.Prometheus
	Conf         *configs.Configs
}

func SetupRoutes(input *SetupRoutesInput) {
	healthRoutes := routes.NewHealthCheck()
	metricsRoutes := routes.NewMetrics()
	customerRoutes := presentation.MakeCustomerRouter(input.ConnPostgres, input.Log, input.Prometheus)
	authRoutes := presentation.MakeAuthRouter(input.ConnPostgres, input.Log, input.Prometheus, input.Conf)

	registerRoutes := NewRegisterRoutes(
		input.Mux,
		input.Log,
		healthRoutes,
		metricsRoutes,
		customerRoutes,
		authRoutes,
	)
	registerRoutes.Register()
}
