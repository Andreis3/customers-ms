package customer

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/command"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/commons"
	"github.com/andreis3/customers-ms/internal/presentation/dtos/input"
	"github.com/andreis3/customers-ms/internal/presentation/dtos/output"
	"github.com/andreis3/customers-ms/internal/presentation/http/transport"
)

type CreateCustomerHandler struct {
	command    command.CreateCustomer
	log        commons.Logger
	prometheus adapter.Prometheus
	tracer     adapter.Tracer
}

func NewCreateCustomerHandler(
	cmd command.CreateCustomer,
	prometheus adapter.Prometheus,
	log commons.Logger,
	tracer adapter.Tracer,
) CreateCustomerHandler {
	return CreateCustomerHandler{
		command:    cmd,
		log:        log,
		prometheus: prometheus,
		tracer:     tracer,
	}
}

func (h *CreateCustomerHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx, span := h.tracer.Start(r.Context(), "CreateCustomerHandler.Handle")
	defer span.End()
	start := time.Now()
	traceID := span.SpanContext().TraceID()

	data, err := transport.DecoderBodyRequest[input.CreatedCustomerDTO](r)
	if err != nil {
		span.RecordError(err)
		h.log.ErrorJSON("failed decode request body",
			slog.String("trace_id", traceID),
			slog.Any("error", err))
		transport.ResponseError[any](w, err)
		return
	}

	res, err := h.command.Execute(ctx, data.MapperToAggregate())
	end := time.Since(start)
	h.log.InfoJSON("end request", slog.String("trace_id", traceID), slog.Float64("duration", float64(end.Milliseconds())))
	if err != nil {
		span.RecordError(err)
		h.log.ErrorJSON("failed execute create customer command",
			slog.String("trace_id", traceID),
			slog.Any("error", err))
		transport.ResponseError[any](w, err)
		return
	}

	h.prometheus.ObserveRequestDuration("/customers", "http", http.StatusCreated, float64(end.Milliseconds()))
	transport.ResponseSuccess(w, http.StatusCreated, output.CustomerOutputMapper(*res))
}
