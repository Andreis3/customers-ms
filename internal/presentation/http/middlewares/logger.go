package middlewares

import (
	"log/slog"
	"net/http"

	"github.com/andreis3/users-ms/internal/domain/interfaces"
	"github.com/andreis3/users-ms/internal/infra/adapters/observability"
)

func LoggingMiddleware(logger interfaces.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, span := observability.Tracer.Start(r.Context(), "HTTP "+r.Method+" "+r.URL.Path)
			defer span.End()
			log := logger.WithTrace(ctx)

			log.Info("new request received",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
			)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
