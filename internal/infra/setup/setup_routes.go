package setup

import (
	"github.com/andreis3/users-ms/internal/app/interfaces"
	"github.com/andreis3/users-ms/internal/infra/routes"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(mux *chi.Mux, postgres interfaces.DB, log interfaces.Logger) {
	routes := routes.NewRegisterRoutes(mux, log)
	routes.Register()
}
