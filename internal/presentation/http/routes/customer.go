package routes

import (
	"net/http"

	"github.com/andreis3/customers-ms/internal/domain/interfaces/adapter"
	"github.com/andreis3/customers-ms/internal/presentation/http/handler"
	"github.com/andreis3/customers-ms/internal/presentation/http/helpers"
	"github.com/andreis3/customers-ms/internal/presentation/http/middlewares"
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

func (r *CustomerRoutes) Routes() helpers.RouteType {
	prefix := "/v1/api/customers"
	return helpers.WithPrefix(prefix, helpers.RouteType{
		{
			Method:      http.MethodPost,
			Handler:     helpers.TraceHandler(http.MethodPost, prefix, r.createCustomer.Handle),
			Description: "Create Customer",
			Middlewares: helpers.Middlewares{
				middlewares.LoggingMiddleware(r.log, r.tracer),
			},
		},
	})
}
