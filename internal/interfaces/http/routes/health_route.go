package routes

import (
	"net/http"

	"github.com/andreis3/users-ms/internal/interfaces/http/handler/health"
	"github.com/andreis3/users-ms/internal/interfaces/http/helpers"
)

type HealthCheckRouter struct{}

func NewHealthCheckRoutes() *HealthCheckRouter {
	return &HealthCheckRouter{}
}

func (r *HealthCheckRouter) HealthCheckRoutes() helpers.RouteType {
	return helpers.RouteType{
		{
			Method:      http.MethodGet,
			Path:        "/health",
			Handler:     health.HealthCheck,
			Description: "Health Check",
			Type:        helpers.HandlerFunc,
			Middlewares: []func(http.Handler) http.Handler{},
		},
	}
}
