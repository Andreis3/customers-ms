package customer

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/andreis3/users-ms/internal/domain/interfaces"
	"github.com/andreis3/users-ms/internal/infra/adapters/observability"
	"github.com/andreis3/users-ms/internal/presentation/http/helpers"
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
	_, child := observability.Tracer.Start(r.Context(), "CreateCustomerHandler.Handle")
	defer child.End()
	start := time.Now()

	tracerID := child.SpanContext().TraceID().String()
	handler.log.InfoJSON("new request received", slog.String("trace_id", tracerID))
	data := struct {
		Status  string `json:"status"`
		TraceID string `json:"trace_id"`
	}{
		Status:  "success",
		TraceID: tracerID,
	}

	end := time.Since(start)
	handler.prometheus.ObserveRequestDuration("/customers", "http", http.StatusCreated, float64(end.Milliseconds()))
	helpers.ResponseSuccess[any](w, http.StatusCreated, data)
}
