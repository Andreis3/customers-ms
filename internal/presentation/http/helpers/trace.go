package helpers

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func TraceHandler(path string, h http.HandlerFunc) http.Handler {
	return otelhttp.NewHandler(http.HandlerFunc(h), path)
}
