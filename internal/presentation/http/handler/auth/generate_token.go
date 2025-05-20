package auth

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/andreis3/customers-ms/internal/app/commands"
	"github.com/andreis3/customers-ms/internal/domain/interfaces"
	"github.com/andreis3/customers-ms/internal/infra/adapters/observability"
	"github.com/andreis3/customers-ms/internal/infra/factories/app"
	"github.com/andreis3/customers-ms/internal/presentation/dtos/output"
	"github.com/andreis3/customers-ms/internal/presentation/http/helpers"
)

type GenerateTokenHandler struct {
	log        interfaces.Logger
	prometheus interfaces.Prometheus
	factory    app.IAuthenticateCustomerFactory
}

func NewGenerateTokenHandler(
	log interfaces.Logger,
	prometheus interfaces.Prometheus,
	customerRepository interfaces.CustomerRepository,
	authService interfaces.Auth,
	bcrypt interfaces.Bcrypt,
) GenerateTokenHandler {
	return GenerateTokenHandler{
		log:        log,
		prometheus: prometheus,
		factory:    app.NewAuthenticateCustomerFactory(log, customerRepository, authService, bcrypt),
	}
}

func (h *GenerateTokenHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx, child := observability.Tracer.Start(r.Context(), "GenerateTokenHandler.Handle")
	start := time.Now()
	traceID := child.SpanContext().TraceID().String()
	defer child.End()

	data, err := helpers.DecoderBodyRequest[commands.AuthenticateCustomerInput](r)
	if err != nil {
		child.RecordError(err)
		h.log.ErrorJSON("failed decode request body",
			slog.String("trace_id", traceID),
			slog.Any("error", err))
		helpers.ResponseError[any](w, err)
		return
	}

	cmd := h.factory.Build()

	input := commands.AuthenticateCustomerInput{
		Email:    data.Email,
		Password: data.Password,
	}

	res, err := cmd.Execute(ctx, input)
	end := time.Since(start)
	if err != nil {
		child.RecordError(err)
		h.log.ErrorJSON("failed execute create customer command",
			slog.String("trace_id", traceID),
			slog.Any("error", err))
		h.log.InfoJSON("end request", slog.String("trace_id", traceID), slog.Float64("duration", float64(end.Milliseconds())))
		helpers.ResponseError[any](w, err)
		return
	}

	h.prometheus.ObserveRequestDuration("/token", "http", http.StatusCreated, float64(end.Milliseconds()))
	h.log.InfoJSON("end request", slog.String("trace_id", traceID), slog.Float64("duration", float64(end.Milliseconds())))
	helpers.ResponseSuccess(w, http.StatusCreated, output.TokenOutputMapper(res))
}
