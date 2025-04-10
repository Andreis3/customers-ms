package routes

import (
	"net/http"

	"github.com/andreis3/users-ms/internal/presentation/http/handler/customer"
	"github.com/andreis3/users-ms/internal/presentation/http/helpers"
)

type CustomerRoutes struct {
	createCustomer customer.CreateCustomerHandler
}

func NewCustomerRoutes(
	createCustomer customer.CreateCustomerHandler,
) *CustomerRoutes {
	return &CustomerRoutes{
		createCustomer: createCustomer,
	}
}

func (r *CustomerRoutes) CustomerRoutes() helpers.RouteType {
	return helpers.RouteType{
		{
			Method:      http.MethodPost,
			Path:        "/customers",
			Handler:     r.createCustomer.Handle,
			Description: "Create Customer",
			Type:        helpers.HandlerFunc,
			Middlewares: []func(http.Handler) http.Handler{},
		},
	}
}
