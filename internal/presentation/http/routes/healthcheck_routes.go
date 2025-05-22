package routes

import (
	"net/http"

	"github.com/andreis3/customers-ms/internal/presentation/http/handler/health"
	"github.com/andreis3/customers-ms/internal/presentation/http/helpers"
)

type HealthCheck struct{}

func NewHealthCheck() *HealthCheck {
	return &HealthCheck{}
}

func (r *HealthCheck) Routes() helpers.RouteType {
	return helpers.RouteType{
		{
			Method:      http.MethodGet,
			Path:        "/health",
			Handler:     health.HealthCheck(),
			Description: "Health Check",
			Middlewares: []func(http.Handler) http.Handler{},
		},
	}
}
