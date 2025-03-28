package routes

import (
	"fmt"
	"net/http"

	"github.com/andreis3/users-ms/internal/app/interfaces"
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

func (r *RegisterRoutes) Register() {}

func (r *RegisterRoutes) registerRoutes(mux *chi.Mux, routesType RouteType) {
	message, info := "[RegisterRoutes] ", "MAPPED_ROUTES"

	for _, route := range routesType {
		switch route.Type {
		case Handler:
			switch len(route.Middlewares) > 0 {
			case true:
				methodAndPath := fmt.Sprintf("%s - %s", route.Method, route.Path)
				r.log.InfoText(message, info, fmt.Sprintf("%s - %s", methodAndPath, route.Description))
				mux.With(route.Middlewares...).Handle(methodAndPath, route.Handler.(http.Handler))
			default:
				methodAndPath := fmt.Sprintf("%s - %s", route.Method, route.Path)
				r.log.InfoText(message, info, fmt.Sprintf("%s - %s", methodAndPath, route.Description))
				mux.Handle(methodAndPath, route.Handler.(http.Handler))
			}
		case HandlerFunc:
			switch len(route.Middlewares) > 0 {
			case true:
				methodAndPath := fmt.Sprintf("%s - %s", route.Method, route.Path)
				r.log.InfoText(message, info, fmt.Sprintf("%s - %s", methodAndPath, route.Description))
				mux.With(route.Middlewares...).HandleFunc(methodAndPath, route.Handler.(func(http.ResponseWriter, *http.Request)))
			default:
				methodAndPath := fmt.Sprintf("%s - %s", route.Method, route.Path)
				r.log.InfoText(message, info, fmt.Sprintf("%s - %s", methodAndPath, route.Description))
				mux.HandleFunc(methodAndPath, route.Handler.(func(http.ResponseWriter, *http.Request)))
			}
		}
	}
}
