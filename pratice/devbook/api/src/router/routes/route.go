package routes

import (
	"api/src/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI      string
	Method   string
	Handler  func(w http.ResponseWriter, r *http.Request)
	NeedAuth bool
}

func Configure(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoutes...)
	routes = append(routes, postsRoutes...)

	for _, route := range routes {
		if route.NeedAuth {
			r.HandleFunc(route.URI, middleware.Logger(middleware.Authenticate(route.Handler))).Methods(route.Method)
		}
		r.HandleFunc(route.URI, middleware.Logger(route.Handler)).Methods(route.Method)
	}
	return r
}
