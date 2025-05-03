package routes

import (
	"github.com/andreis3/users-ms/internal/domain/interfaces"
	"github.com/andreis3/users-ms/internal/infra/adapters/crypto"
	"github.com/andreis3/users-ms/internal/infra/uow"
	"github.com/andreis3/users-ms/internal/presentation/http/handler/customer"
	"github.com/andreis3/users-ms/internal/presentation/http/routes"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRoutes(mux *chi.Mux, postgres interfaces.DB, log interfaces.Logger, prometheus interfaces.Prometheus) {
	postgresPool := postgres.Instance().(*pgxpool.Pool)
	crypto := crypto.NewBcrypt()
	uow := uow.NewUnitOfWork(postgresPool, prometheus)
	createCustomerHandler := customer.NewCreateCustomerHandler(log, prometheus, crypto, uow)
	customerRoutes := routes.NewCustomerRoutes(createCustomerHandler, log)
	routes := NewRegisterRoutes(mux, log, *customerRoutes)
	routes.Register()
}
