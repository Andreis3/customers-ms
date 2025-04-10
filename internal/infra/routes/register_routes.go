package routes

import (
	"fmt"
	"net/http"

	"github.com/andreis3/users-ms/internal/domain/interfaces"
	"github.com/andreis3/users-ms/internal/presentation/http/helpers"
	"github.com/andreis3/users-ms/internal/presentation/http/routes"
	"github.com/go-chi/chi/v5"
)

type RegisterRoutes struct {
	mux            *chi.Mux
	log            interfaces.Logger
	customerRoutes routes.CustomerRoutes
}

func NewRegisterRoutes(
	mux *chi.Mux,
	log interfaces.Logger,
	customerRoutes routes.CustomerRoutes,
) *RegisterRoutes {
	return &RegisterRoutes{
		mux:            mux,
		log:            log,
		customerRoutes: customerRoutes,
	}
}

func (r *RegisterRoutes) Register() {
	// Example: here you register the HealthCheck routes;
	// For other routes, just call them the same way.
	r.registerRoutes(routes.NewHealthCheck().HealthCheck())
	r.registerRoutes(routes.NewMetrics().Metrics())
	r.registerRoutes(r.customerRoutes.CustomerRoutes())
}

// registerRoutes iterates over the returned routes
// and calls attachRoute for each one.
func (r *RegisterRoutes) registerRoutes(routeDefs helpers.RouteType) {
	for _, route := range routeDefs {
		r.attachRoute(route)
	}
}

// attachRoute encapsulates the logic of:
// 1) Logging method and path,
// 2) Applying middlewares (if any),
// 3) Registering the handler correctly.
func (r *RegisterRoutes) attachRoute(route helpers.RouteFields) {
	methodAndPath := fmt.Sprintf("%s %s", route.Method, route.Path)
	r.log.InfoText("[RegisterRoutes] ", "MAPPED_ROUTES", fmt.Sprintf("%s - %s", methodAndPath, route.Description))

	// If middlewares exist, we apply them via .With(...)
	// and register within a .Group
	if len(route.Middlewares) > 0 {
		r.mux.With(route.Middlewares...).Group(func(m chi.Router) {
			r.registerHandler(m, route)
		})
	} else {
		// Without middlewares, we register directly
		r.registerHandler(r.mux, route)
	}
}

// registerHandler checks whether route.Handler is a Handler
func (r *RegisterRoutes) registerHandler(m chi.Router, route helpers.RouteFields) {
	handler, ok := route.Handler.(http.Handler)
	if !ok {
		r.log.CriticalText("Route registration error: invalid handler type for Handler")
		return
	}

	// Method(...) to explicitly register the HTTP method
	m.Method(route.Method, route.Path, handler)
}
