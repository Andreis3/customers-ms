package routes

import (
	"net/http"

	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/service"
	"github.com/andreis3/customers-ms/internal/presentation/http/handler"
	"github.com/andreis3/customers-ms/internal/presentation/http/helpers"
	"github.com/andreis3/customers-ms/internal/presentation/http/middlewares"
)

type AddressRoutes struct {
	getAddress  handler.GetAddressHandler
	authService service.Auth
	log         adapter.Logger
	tracer      adapter.Tracer
}

func NewAddressRoutes(
	getAddress handler.GetAddressHandler,
	authService service.Auth,
	log adapter.Logger,
	tracer adapter.Tracer,
) *AddressRoutes {
	return &AddressRoutes{
		getAddress:  getAddress,
		authService: authService,
		log:         log,
		tracer:      tracer,
	}
}

func (r *AddressRoutes) Routes() helpers.RouteType {
	prefix := "/v1/api/address"
	return helpers.WithPrefix(prefix, helpers.RouteType{
		{
			Method:      http.MethodGet,
			Handler:     helpers.TraceHandler(http.MethodGet, prefix, r.getAddress.Handle),
			Description: "Get Address",
			Middlewares: helpers.Middlewares{
				middlewares.LoggingMiddleware(r.log, r.tracer),
				middlewares.ValidateCustomer(r.authService, r.log, r.tracer),
			},
		},
	})
}
