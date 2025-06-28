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

func SetupRoutes(mux *chi.Mux, connPostgres *postegres.Postgres, log commons.Logger, prometheus adapter.Prometheus, conf *configs.Configs) {
	healthRoutes := routes.NewHealthCheck()
	metricsRoutes := routes.NewMetrics()
	customerRoutes := presentation.MakeCustomerRouter(connPostgres, log, prometheus)
	authRoutes := presentation.MakeAuthRouter(connPostgres, log, prometheus, conf)

	registerRoutes := NewRegisterRoutes(
		mux,
		log,
		healthRoutes,
		metricsRoutes,
		customerRoutes,
		authRoutes,
	)
	registerRoutes.Register()
}
