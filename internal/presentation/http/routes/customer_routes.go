package routes

import (
	"net/http"

	"github.com/andreis3/customers-ms/internal/domain/interfaces"
	"github.com/andreis3/customers-ms/internal/presentation/http/handler/customer"
	"github.com/andreis3/customers-ms/internal/presentation/http/helpers"
	"github.com/andreis3/customers-ms/internal/presentation/http/middlewares"
)

type CustomerRoutes struct {
	createCustomer customer.CreateCustomerHandler
	log            interfaces.Logger
}

func NewCustomerRoutes(
	createCustomer customer.CreateCustomerHandler,
	log interfaces.Logger,
) *CustomerRoutes {
	return &CustomerRoutes{
		createCustomer: createCustomer,
		log:            log,
	}
}

func (r *CustomerRoutes) Routes() helpers.RouteType {
	prefix := "/v1/api/customers"
	return helpers.WithPrefix(prefix, helpers.RouteType{
		{
			Method:      http.MethodPost,
			Handler:     helpers.TraceHandler(http.MethodPost, prefix, r.createCustomer.Handle),
			Description: "Create Customer",
			Middlewares: []func(http.Handler) http.Handler{
				middlewares.LoggingMiddleware(r.log),
			},
		},
	})
}
