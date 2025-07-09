package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/infra/adapters/db"
	"github.com/andreis3/customers-ms/internal/infra/configs"
	"github.com/andreis3/customers-ms/internal/infra/factories/presentation"
	"github.com/andreis3/customers-ms/internal/presentation/http/routes"
)

type RegisterRoutesDeps struct {
	Mux        *chi.Mux
	PostgresDB *db.Postgres
	Log        adapter.Logger
	Prometheus adapter.Prometheus
	Conf       *configs.Configs
	Tracer     adapter.Tracer
}

func Setup(deps *RegisterRoutesDeps) {
	registerRoutes := NewRegisterRoutes(
		deps.Mux,
		deps.Log,
		BuildRoutes(deps),
	)

	registerRoutes.Register()
}

func BuildRoutes(deps *RegisterRoutesDeps) []ModuleRoutes {
	return []ModuleRoutes{
		routes.NewHealthCheck(),
		routes.NewMetrics(),
		presentation.MakeCustomerRouter(deps.PostgresDB, deps.Log, deps.Prometheus, deps.Tracer),
		presentation.MakeAuthRouter(deps.PostgresDB, deps.Log, deps.Prometheus, deps.Conf, deps.Tracer),
	}
}
