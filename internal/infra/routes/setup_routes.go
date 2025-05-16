package routes

import (
	"github.com/andreis3/customers-ms/internal/domain/interfaces"
	"github.com/andreis3/customers-ms/internal/domain/services"
	"github.com/andreis3/customers-ms/internal/infra/adapters/crypto"
	"github.com/andreis3/customers-ms/internal/infra/adapters/db/postegres"
	"github.com/andreis3/customers-ms/internal/infra/repositories/postgres/repository"
	"github.com/andreis3/customers-ms/internal/infra/uow"
	"github.com/andreis3/customers-ms/internal/presentation/http/handler/customer"
	"github.com/andreis3/customers-ms/internal/presentation/http/routes"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(mux *chi.Mux, connPostgres *postegres.Postgres, log interfaces.Logger, prometheus interfaces.Prometheus) {
	pool := connPostgres.Pool
	crypto := crypto.NewBcrypt()
	uow := uow.NewUnitOfWork(pool, prometheus)
	customerRepository := repository.NewCustomerRepository(pool, prometheus)
	customerService := services.NewCustomerService(customerRepository)
	createCustomerHandler := customer.NewCreateCustomerHandler(log, prometheus, crypto, uow, customerService)

	// Routes
	customerRoutes := routes.NewCustomerRoutes(createCustomerHandler, log)
	healthRoutes := routes.NewHealthCheck()
	metricsRoutes := routes.NewMetrics()

	routes := NewRegisterRoutes(
		mux,
		log,
		healthRoutes,
		metricsRoutes,
		customerRoutes,
	)
	routes.Register()
}
