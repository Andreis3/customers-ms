package handler

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/andreis3/customers-ms/internal/app/dto"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/command"
	"github.com/andreis3/customers-ms/internal/presentation/http/transport"
)

type GenerateTokenHandler struct {
	log        adapter.Logger
	prometheus adapter.Prometheus
	cmd        command.Login
	tracer     adapter.Tracer
}

func NewGenerateTokenHandler(
	log adapter.Logger,
	prometheus adapter.Prometheus,
	tracer adapter.Tracer,
	cmd command.Login,
) GenerateTokenHandler {
	return GenerateTokenHandler{
		log:        log,
		prometheus: prometheus,
		cmd:        cmd,
		tracer:     tracer,
	}
}

func (h *GenerateTokenHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx, span := h.tracer.Start(r.Context(), "GenerateTokenHandler.Handle")
	start := time.Now()
	traceID := span.SpanContext().TraceID()
	defer span.End()

	input, err := transport.DecoderBodyRequest[dto.LoginInput](r)
	if err != nil {
		span.RecordError(err)
		h.log.ErrorJSON("failed decode request body",
			slog.String("trace_id", traceID),
			slog.Any("error", err))
		transport.ResponseError(w, err)
		return
	}

	res, err := h.cmd.Execute(ctx, input)
	end := time.Since(start)
	if err != nil {
		span.RecordError(err)
		h.log.ErrorJSON("failed execute create customer command",
			slog.String("trace_id", traceID),
			slog.Any("error", err))
		h.log.InfoJSON("end request", slog.String("trace_id", traceID), slog.Float64("duration", float64(end.Milliseconds())))
		transport.ResponseError(w, err)
		return
	}

	h.prometheus.ObserveRequestDuration("/token", "http", http.StatusCreated, float64(end.Milliseconds()))
	h.log.InfoJSON("end request", slog.String("trace_id", traceID), slog.Float64("duration", float64(end.Milliseconds())))
	transport.ResponseSuccess(w, http.StatusCreated, res)
}
