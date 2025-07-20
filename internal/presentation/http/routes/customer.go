package routes

import (
	"net/http"

	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/service"
	"github.com/andreis3/customers-ms/internal/presentation/http/handler"
	"github.com/andreis3/customers-ms/internal/presentation/http/helpers"
	"github.com/andreis3/customers-ms/internal/presentation/http/middlewares"
)

type CustomerRoutes struct {
	createCustomer *handler.CreateCustomerHandler
	getAddress     *handler.GetCustomerAddressesHandler
	authService    service.Auth
	log            adapter.Logger
	tracer         adapter.Tracer
}

func NewCustomer(
	createCustomer *handler.CreateCustomerHandler,
	getAddress *handler.GetCustomerAddressesHandler,
	authService service.Auth,
	log adapter.Logger,
	tracer adapter.Tracer,
) *CustomerRoutes {
	return &CustomerRoutes{
		createCustomer: createCustomer,
		getAddress:     getAddress,
		authService:    authService,
		log:            log,
		tracer:         tracer,
	}
}

func (r *CustomerRoutes) Routes() helpers.RouteType {
	prefix := "/v1/api"
	return helpers.WithPrefix(prefix, helpers.RouteType{
		{
			Method:      http.MethodPost,
			Path:        "/customer",
			Handler:     helpers.TraceHandler(http.MethodPost, prefix+"/customer", r.createCustomer.Handle),
			Description: "Create Customer",
			Middlewares: helpers.Middlewares{
				middlewares.LoggingMiddleware(r.log, r.tracer),
			},
		},
		{
			Method:      http.MethodGet,
			Path:        "/customer/addresses",
			Handler:     helpers.TraceHandler(http.MethodGet, prefix+"/customer/addresses", r.getAddress.Handle),
			Description: "Get Customer Addresses",
			Middlewares: helpers.Middlewares{
				middlewares.LoggingMiddleware(r.log, r.tracer),
				middlewares.ValidateCustomer(r.authService, r.log, r.tracer),
			},
		},
	})
}
