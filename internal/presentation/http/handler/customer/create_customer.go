package customer

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/andreis3/users-ms/internal/domain/interfaces"
	"github.com/andreis3/users-ms/internal/presentation/http/helpers"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type CreateCustomerHandler struct {
	log        interfaces.Logger
	prometheus interfaces.Prometheus
}

func NewCreateCustomerHandler(
	log interfaces.Logger,
	prometheus interfaces.Prometheus,
) CreateCustomerHandler {
	return CreateCustomerHandler{
		log:        log,
		prometheus: prometheus,
	}
}

func (handler *CreateCustomerHandler) Handle(w http.ResponseWriter, r *http.Request) {
	span := trace.SpanFromContext(r.Context())
	defer span.End()
	span.AddEvent("Start execute handler", trace.WithAttributes())
	start := time.Now()

	tracerID := span.SpanContext().TraceID().String()
	handler.log.InfoJSON("new request received", slog.String("trace_id", tracerID))
	data := struct {
		Status string `json:"status"`
	}{
		Status: "success",
	}

	end := time.Since(start)
	handler.prometheus.ObserveRequestDuration("/customers", "http", http.StatusCreated, float64(end.Milliseconds()))
	span.AddEvent("Request processed", trace.WithAttributes(
		attribute.String("handler", "CreateCustomerHandler"),
		attribute.Float64("duration_ms", float64(end.Microseconds())),
	))
	helpers.ResponseSuccess[any](w, http.StatusCreated, data)
}
