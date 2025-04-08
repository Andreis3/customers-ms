package observability

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	api "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric"
)

const (
	MeterName    = "customers-ms"
	MeterVersion = "1.0.0"
)

type Prometheus struct {
	counterRequestStatusCode api.Int64Counter
}

func NewPrometheus() *Prometheus {
	exporter, _ := prometheus.New()
	provider := metric.NewMeterProvider(metric.WithReader(exporter))
	meter := provider.Meter(MeterName, api.WithInstrumentationVersion(MeterVersion))

	counterRequestStatusCode, _ := meter.Int64Counter("proxy_requests_total",
		api.WithDescription("Total number of proxy requests"))

	return &Prometheus{counterRequestStatusCode: counterRequestStatusCode}
}

func (p *Prometheus) CounterRequestStatusCode(router, protocol string, statusCode int) {
	opt := api.WithAttributes(
		attribute.Key("router").String(router),
		attribute.Key("status_code").Int(statusCode),
		attribute.Key("protocol").String(protocol),
	)
	p.counterRequestStatusCode.Add(context.Background(), 1, opt)
}
