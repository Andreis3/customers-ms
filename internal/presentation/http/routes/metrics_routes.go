package routes

import (
	"net/http"

	"github.com/andreis3/customers-ms/internal/presentation/http/transport"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metrics struct{}

func NewMetrics() *Metrics {
	return &Metrics{}
}

func (m *Metrics) Routes() transport.RouteType {
	return transport.RouteType{
		{
			Method:      http.MethodGet,
			Path:        "/metrics",
			Handler:     promhttp.Handler(),
			Description: "Metrics Prometheus",
			Middlewares: []func(http.Handler) http.Handler{},
		},
	}
}
