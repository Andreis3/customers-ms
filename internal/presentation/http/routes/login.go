package routes

import (
	"net/http"

	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/presentation/http/handler"
	"github.com/andreis3/customers-ms/internal/presentation/http/middlewares"
	"github.com/andreis3/customers-ms/internal/presentation/http/transport"
)

type LoginRoutes struct {
	log         adapter.Logger
	authHandler handler.LoginCustomer
	tracer      adapter.Tracer
}

func NewLoginRoutes(
	log adapter.Logger,
	authHandler handler.LoginCustomer,
	tracer adapter.Tracer,
) *LoginRoutes {
	return &LoginRoutes{
		log:         log,
		authHandler: authHandler,
		tracer:      tracer,
	}
}

func (r *LoginRoutes) Routes() transport.RouteType {
	prefix := "/v1/api/login"
	return transport.WithPrefix(prefix, transport.RouteType{
		{
			Method:      http.MethodPost,
			Handler:     transport.TraceHandler(http.MethodPost, prefix, r.authHandler.Handle),
			Description: "Generate Token",
			Middlewares: []func(http.Handler) http.Handler{
				middlewares.LoggingMiddleware(r.log, r.tracer),
			},
		},
	})
}
