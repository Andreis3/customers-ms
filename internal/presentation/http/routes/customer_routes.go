package routes

import (
	"net/http"

	"github.com/andreis3/users-ms/internal/domain/interfaces"
	"github.com/andreis3/users-ms/internal/presentation/http/handler/customer"
	"github.com/andreis3/users-ms/internal/presentation/http/helpers"
	"github.com/andreis3/users-ms/internal/presentation/http/middlewares"
)

const (
	CustomersPath = "/customers"
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

func (r *CustomerRoutes) CustomerRoutes() helpers.RouteType {
	return helpers.RouteType{
		{
			Method:      http.MethodPost,
			Path:        CustomersPath,
			Handler:     helpers.TraceHandler(CustomersPath, r.createCustomer.Handle),
			Description: "Create Customer",
			Middlewares: []func(http.Handler) http.Handler{
				middlewares.LoggingMiddleware(r.log),
			},
		},
	}
}
