package routes

import (
	"net/http"

	"github.com/andreis3/users-ms/internal/presentation/http/helpers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metrics struct{}

func NewMetrics() *Metrics {
	return &Metrics{}
}

func (m *Metrics) Metrics() helpers.RouteType {
	return helpers.RouteType{
		{
			Method:      http.MethodGet,
			Path:        "/metrics",
			Handler:     promhttp.Handler(),
			Description: "Metrics Prometheus",
			Middlewares: []func(http.Handler) http.Handler{},
		},
	}
}
