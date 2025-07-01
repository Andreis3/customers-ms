package routes

import (
	"net/http"

	"github.com/andreis3/customers-ms/internal/domain/interfaces/commons"
	"github.com/andreis3/customers-ms/internal/presentation/http/handler/customer"
	"github.com/andreis3/customers-ms/internal/presentation/http/middlewares"
	"github.com/andreis3/customers-ms/internal/presentation/http/transport"
)

type CustomerRoutes struct {
	createCustomer customer.CreateCustomerHandler
	log            commons.Logger
}

func NewCustomer(
	createCustomer customer.CreateCustomerHandler,
	log commons.Logger,
) *CustomerRoutes {
	return &CustomerRoutes{
		createCustomer: createCustomer,
		log:            log,
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
				middlewares.LoggingMiddleware(r.log),
			},
		},
	})
}
