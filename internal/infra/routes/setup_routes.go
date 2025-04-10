package routes

import (
	"github.com/andreis3/users-ms/internal/domain/interfaces"
	"github.com/andreis3/users-ms/internal/infra/adapters/observability"
	"github.com/andreis3/users-ms/internal/presentation/http/handler/customer"
	"github.com/andreis3/users-ms/internal/presentation/http/routes"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(mux *chi.Mux, postgres interfaces.DB, log interfaces.Logger) {
	prometheus := observability.NewPrometheus()
	createCustomerHandler := customer.NewCreateCustomerHandler(log, prometheus)
	customerRoutes := routes.NewCustomerRoutes(createCustomerHandler)
	routes := NewRegisterRoutes(mux, log, *customerRoutes)
	routes.Register()
}
