package routes

import "net/http"

const (
	Handler     = "handler"
	HandlerFunc = "handlerFunc"
)

type RouteType []RouteFields

type RouteFields struct {
	Method      string
	Path        string
	Handler     any
	Description string
	Type        string
	Middlewares []func(http.Handler) http.Handler
}
