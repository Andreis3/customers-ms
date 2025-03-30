package routes

import (
	"github.com/andreis3/users-ms/internal/app/interfaces"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(mux *chi.Mux, postgres interfaces.DB, log interfaces.Logger) {
	routes := NewRegisterRoutes(mux, log)
	routes.Register()
}
