package customer

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/andreis3/users-ms/internal/domain/interfaces"
	"github.com/andreis3/users-ms/internal/infra/adapters/observability"
	"github.com/andreis3/users-ms/internal/infra/factory/command"
	"github.com/andreis3/users-ms/internal/presentation/dtos/input"
	"github.com/andreis3/users-ms/internal/presentation/dtos/output"
	"github.com/andreis3/users-ms/internal/presentation/http/helpers"
)

type CreateCustomerHandler struct {
	log        interfaces.Logger
	prometheus interfaces.Prometheus
	uow        interfaces.UnitOfWork
}

func NewCreateCustomerHandler(
	log interfaces.Logger,
	prometheus interfaces.Prometheus,
	uow interfaces.UnitOfWork,
) CreateCustomerHandler {
	return CreateCustomerHandler{
		log:        log,
		prometheus: prometheus,
		uow:        uow,
	}
}

func (handler *CreateCustomerHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx, child := observability.Tracer.Start(r.Context(), "CreateCustomerHandler.Handle")
	start := time.Now()
	traceID := child.SpanContext().TraceID().String()
	defer child.End()

	data, err := helpers.DecoderBodyRequest[input.CreatedCustomerDTO](r)
	if err != nil {
		child.RecordError(err)
		handler.log.ErrorJSON("failed decode request body",
			slog.String("trace_id", traceID),
			slog.Any("error", err))
		helpers.ResponseError[any](w, err)
		return
	}

	cmd := command.NewCreatedCustomerFactory(handler.uow)
	res, err := cmd.Execute(ctx, data.MapperToAggregate())
	end := time.Since(start)
	handler.log.InfoJSON("end request", slog.String("trace_id", traceID), slog.Float64("duration", float64(end.Milliseconds())))
	if err != nil {
		child.RecordError(err)
		handler.log.ErrorJSON("failed execute create customer command",
			slog.String("trace_id", traceID),
			slog.Any("error", err))
		helpers.ResponseError[any](w, err)
		return
	}

	handler.prometheus.ObserveRequestDuration("/customers", "http", http.StatusCreated, float64(end.Milliseconds()))
	helpers.ResponseSuccess(w, http.StatusCreated, output.CustomerOutputMapper(*res))
}
