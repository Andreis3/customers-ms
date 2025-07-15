package routes

import (
	"net/http"

	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/presentation/http/handler"
	"github.com/andreis3/customers-ms/internal/presentation/http/helpers"
	"github.com/andreis3/customers-ms/internal/presentation/http/middlewares"
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

func (r *LoginRoutes) Routes() helpers.RouteType {
	prefix := "/v1/api/login"
	return helpers.WithPrefix(prefix, helpers.RouteType{
		{
			Method:      http.MethodPost,
			Handler:     helpers.TraceHandler(http.MethodPost, prefix, r.authHandler.Handle),
			Description: "Generate Token",
			Middlewares: helpers.Middlewares{
				middlewares.LoggingMiddleware(r.log, r.tracer),
			},
		},
	})
}
