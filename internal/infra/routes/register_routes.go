package routes

import (
	"fmt"
	"net/http"

	"github.com/andreis3/users-ms/internal/app/interfaces"
	"github.com/andreis3/users-ms/internal/interfaces/http/helpers"
	"github.com/andreis3/users-ms/internal/interfaces/http/routes"
	"github.com/go-chi/chi/v5"
)

type RegisterRoutes struct {
	mux *chi.Mux
	log interfaces.Logger
}

func NewRegisterRoutes(mux *chi.Mux, log interfaces.Logger) *RegisterRoutes {
	return &RegisterRoutes{
		mux: mux,
		log: log,
	}
}

func (r *RegisterRoutes) Register() {

  // Example: here you register the HealthCheck routes;
  // For other routes, just call them the same way.
	r.registerRoutes(routes.NewHealthCheckRoutes().HealthCheckRoutes())
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

// registerHandler checks whether route.Handler is a Handler or HandlerFunc
// and registers it according to the Type defined in helpers.RouteType.
func (r *RegisterRoutes) registerHandler(m chi.Router, route helpers.RouteFields) {
	switch route.Type {
	case helpers.Handler:
		handler, ok := route.Handler.(http.Handler)
		if !ok {
			r.log.ErrorText("Route registration error: invalid handler type for Handler")
			return
		}

		// Method(...) to explicitly register the HTTP method
		m.Method(route.Method, route.Path, handler)

	case helpers.HandlerFunc:
		hf, ok := route.Handler.(func(http.ResponseWriter, *http.Request))
		if !ok {
			r.log.ErrorText("Route registration error: invalid handler type for HandlerFunc")
			return
		}
		// MethodFunc(...) to explicitly register the HTTP method
		m.MethodFunc(route.Method, route.Path, hf)

	default:
		r.log.ErrorText("Route registration error: unknown route type")
	}
}
