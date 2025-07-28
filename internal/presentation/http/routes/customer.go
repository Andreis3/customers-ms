package routes

import (
	"net/http"

	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/domain/interfaces/service"
	"github.com/andreis3/customers-ms/internal/infra/factories/presentation/handler-factory"
	"github.com/andreis3/customers-ms/internal/presentation/http/helpers"
	"github.com/andreis3/customers-ms/internal/presentation/http/middlewares"
)

type CustomerRoutes struct {
	createCustomerHandler *handler_factory.CreateCustomer
	getAddress            *handler_factory.GetCustomerAddresses
	authService           service.Auth
	log                   adapter.Logger
	tracer                adapter.Tracer
}

func NewCustomer(
	createCustomerHandler *handler_factory.CreateCustomer,
	getAddress *handler_factory.GetCustomerAddresses,
	authService service.Auth,
	log adapter.Logger,
	tracer adapter.Tracer,
) *CustomerRoutes {
	return &CustomerRoutes{
		createCustomerHandler: createCustomerHandler,
		getAddress:            getAddress,
		authService:           authService,
		log:                   log,
		tracer:                tracer,
	}
}

func (cr *CustomerRoutes) Routes() helpers.RouteType {
	prefix := "/v1/api"
	return helpers.WithPrefix(prefix, helpers.RouteType{
		{
			Method: http.MethodPost,
			Path:   "/customer",
			Handler: helpers.TraceHandler(http.MethodPost, prefix+"/customer", func(w http.ResponseWriter, r *http.Request) {
				h := cr.createCustomerHandler.NewCreateCustomer()
				h.Handle(w, r)
			}),
			Description: "Create Customer",
			Middlewares: helpers.Middlewares{
				middlewares.LoggingMiddleware(cr.log, cr.tracer),
			},
		},
		{
			Method: http.MethodGet,
			Path:   "/customer/addresses",
			Handler: helpers.TraceHandler(http.MethodGet, prefix+"/customer/addresses", func(w http.ResponseWriter, r *http.Request) {
				h := cr.getAddress.NewGetCustomerAddresses()
				h.Handle(w, r)
			}),
			Description: "Get Customer Addresses",
			Middlewares: helpers.Middlewares{
				middlewares.LoggingMiddleware(cr.log, cr.tracer),
				middlewares.ValidateCustomer(cr.authService, cr.log, cr.tracer),
			},
		},
	})
}
