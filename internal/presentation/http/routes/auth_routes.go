package routes

import (
	"net/http"

	"github.com/andreis3/customers-ms/internal/domain/interfaces/commons"
	"github.com/andreis3/customers-ms/internal/presentation/http/handler/auth"
	"github.com/andreis3/customers-ms/internal/presentation/http/helpers"
	"github.com/andreis3/customers-ms/internal/presentation/http/middlewares"
)

type AuthRoutes struct {
	log         commons.Logger
	authHandler auth.GenerateTokenHandler
}

func NewAuthRoutes(
	log commons.Logger,
	authHandler auth.GenerateTokenHandler,
) *AuthRoutes {
	return &AuthRoutes{
		log:         log,
		authHandler: authHandler,
	}
}

func (r *AuthRoutes) Routes() helpers.RouteType {
	prefix := "/v1/api/auth"
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
