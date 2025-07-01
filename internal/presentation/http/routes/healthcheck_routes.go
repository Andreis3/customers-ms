package routes

import (
	"net/http"

	"github.com/andreis3/customers-ms/internal/presentation/http/handler/health"
	"github.com/andreis3/customers-ms/internal/presentation/http/transport"
)

type HealthCheck struct{}

func NewHealthCheck() *HealthCheck {
	return &HealthCheck{}
}

func (r *HealthCheck) Routes() transport.RouteType {
	return transport.RouteType{
		{
			Method:      http.MethodGet,
			Path:        "/health",
			Handler:     health.HealthCheck(),
			Description: "Health Check",
			Middlewares: []func(http.Handler) http.Handler{},
		},
	}
}
