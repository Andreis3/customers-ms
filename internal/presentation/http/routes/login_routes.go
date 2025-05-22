package routes

import (
	"net/http"

	"github.com/andreis3/customers-ms/internal/domain/interfaces/commons"
	"github.com/andreis3/customers-ms/internal/presentation/http/handler/login"
	"github.com/andreis3/customers-ms/internal/presentation/http/helpers"
	"github.com/andreis3/customers-ms/internal/presentation/http/middlewares"
)

type LoginRoutes struct {
	log         commons.Logger
	authHandler login.GenerateTokenHandler
}

func NewAuthRoutes(
	log commons.Logger,
	authHandler login.GenerateTokenHandler,
) *LoginRoutes {
	return &LoginRoutes{
		log:         log,
		authHandler: authHandler,
	}
}

func (r *LoginRoutes) Routes() helpers.RouteType {
	prefix := "/v1/api/login"
	return helpers.WithPrefix(prefix, helpers.RouteType{
		{
			Method:      http.MethodPost,
			Handler:     helpers.TraceHandler(http.MethodPost, prefix, r.authHandler.Handle),
			Description: "Generate Token",
			Middlewares: []func(http.Handler) http.Handler{
				middlewares.LoggingMiddleware(r.log),
			},
		},
	})
}
