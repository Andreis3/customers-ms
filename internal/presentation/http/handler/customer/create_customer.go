package customer

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/andreis3/users-ms/internal/domain/interfaces"
	"github.com/andreis3/users-ms/internal/infra/adapters/observability"
	"github.com/andreis3/users-ms/internal/infra/factories"
	"github.com/andreis3/users-ms/internal/presentation/dtos/input"
	"github.com/andreis3/users-ms/internal/presentation/dtos/output"
	"github.com/andreis3/users-ms/internal/presentation/http/helpers"
)

type CreateCustomerHandler struct {
	log        interfaces.Logger
	prometheus interfaces.Prometheus
	factory    factories.ICreateCustomerFactory
}

func NewCreateCustomerHandler(
	log interfaces.Logger,
	prometheus interfaces.Prometheus,
	crypto interfaces.Bcrypt,
	uow interfaces.UnitOfWork,
) CreateCustomerHandler {
	return CreateCustomerHandler{
		log:        log,
		prometheus: prometheus,
		factory:    factories.NewCreateCustomerFactory(uow, crypto, log),
	}
}

func (h *CreateCustomerHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx, child := observability.Tracer.Start(r.Context(), "CreateCustomerHandler.Handle")
	start := time.Now()
	traceID := child.SpanContext().TraceID().String()
	defer child.End()

	data, err := helpers.DecoderBodyRequest[input.CreatedCustomerDTO](r)
	if err != nil {
		child.RecordError(err)
		h.log.ErrorJSON("failed decode request body",
			slog.String("trace_id", traceID),
			slog.Any("error", err))
		helpers.ResponseError[any](w, err)
		return
	}

	cmd := h.factory.Build()
	res, err := cmd.Execute(ctx, data.MapperToAggregate())
	end := time.Since(start)
	h.log.InfoJSON("end request", slog.String("trace_id", traceID), slog.Float64("duration", float64(end.Milliseconds())))
	if err != nil {
		child.RecordError(err)
		h.log.ErrorJSON("failed execute create customer command",
			slog.String("trace_id", traceID),
			slog.Any("error", err))
		helpers.ResponseError[any](w, err)
		return
	}

	h.prometheus.ObserveRequestDuration("/customers", "http", http.StatusCreated, float64(end.Milliseconds()))
	helpers.ResponseSuccess(w, http.StatusCreated, output.CustomerOutputMapper(*res))
}
