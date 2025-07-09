package routes

import (
	"net/http"

	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/presentation/http/handler"
	"github.com/andreis3/customers-ms/internal/presentation/http/middlewares"
	"github.com/andreis3/customers-ms/internal/presentation/http/transport"
)

type CustomerRoutes struct {
	createCustomer handler.CreateCustomerHandler
	log            adapter.Logger
	tracer         adapter.Tracer
}

func NewCustomer(
	createCustomer handler.CreateCustomerHandler,
	log adapter.Logger,
	tracer adapter.Tracer,
) *CustomerRoutes {
	return &CustomerRoutes{
		createCustomer: createCustomer,
		log:            log,
		tracer:         tracer,
	}
}

func (r *CustomerRoutes) Routes() transport.RouteType {
	prefix := "/v1/api/customers"
	return transport.WithPrefix(prefix, transport.RouteType{
		{
			Method:      http.MethodPost,
			Handler:     transport.TraceHandler(http.MethodPost, prefix, r.createCustomer.Handle),
			Description: "Create Customer",
			Middlewares: []func(http.Handler) http.Handler{
				middlewares.LoggingMiddleware(r.log, r.tracer),
			},
		},
	})
}
