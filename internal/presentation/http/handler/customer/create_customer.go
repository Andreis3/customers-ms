package customer

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/commons"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/service"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/uow"
	"github.com/andreis3/customers-ms/internal/infra/adapters/observability"
	"github.com/andreis3/customers-ms/internal/infra/factories/app"
	"github.com/andreis3/customers-ms/internal/presentation/dtos/input"
	"github.com/andreis3/customers-ms/internal/presentation/dtos/output"
	"github.com/andreis3/customers-ms/internal/presentation/http/helpers"
)

type CreateCustomerHandler struct {
	log        commons.Logger
	prometheus adapter.Prometheus
	factory    app.CreateCustomerFactory
}

func NewCreateCustomerHandler(
	log commons.Logger,
	prometheus adapter.Prometheus,
	crypto adapter.Bcrypt,
	uow uow.UnitOfWork,
	customerService service.CustomerService,
) CreateCustomerHandler {
	return CreateCustomerHandler{
		log:        log,
		prometheus: prometheus,
		factory:    app.NewCreateCustomerFactory(uow, crypto, log, customerService),
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
