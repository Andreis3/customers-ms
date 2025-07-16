package middlewares

import (
	"log/slog"
	"net/http"

	"github.com/andreis3/customers-ms/internal/domain/errors"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/service"
	"github.com/andreis3/customers-ms/internal/presentation/http/helpers"
)

func ValidateCustomer(authService service.Auth, logger adapter.Logger, tracer adapter.Tracer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, span := tracer.Start(r.Context(), "ValidateCustomer")
			defer span.End()
			traceID := span.SpanContext().TraceID()

			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				logger.ErrorJSON("missing authorization header",
					slog.String("trace_id", traceID),
					slog.String("method", r.Method),
					slog.String("path", r.URL.Path))
				helpers.ResponseError(w, errors.ErrorInvalidToken())
				return
			}

			token := authHeader[len("Bearer "):]
			if token == "" {
				logger.ErrorJSON("missing token",
					slog.String("trace_id", traceID),
					slog.String("method", r.Method),
					slog.String("path", r.URL.Path))
				helpers.ResponseError(w, errors.ErrorInvalidToken())
				return
			}

			_, err := authService.DecodeToken(ctx, token)
			if err != nil {
				logger.ErrorJSON("failed validate token",
					slog.String("trace_id", traceID),
					slog.String("method", r.Method),
					slog.String("path", r.URL.Path),
					slog.Any("error", err))
				helpers.ResponseError(w, errors.ErrorInvalidToken())
				return
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
