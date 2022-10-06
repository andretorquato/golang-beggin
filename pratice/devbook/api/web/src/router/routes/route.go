package routes

import (
	"net/http"
	"web/src/middleware"

	"github.com/gorilla/mux"
)

type Route struct {
	URI      string
	Method   string
	Handle   func(w http.ResponseWriter, r *http.Request)
	NeedAuth bool
}

func Configure(router *mux.Router) *mux.Router {
	routes := loginRoutes
	routes = append(routes, usersRoutes...)
	routes = append(routes, homeRoutes...)

	for _, route := range routes {
		if route.NeedAuth {
			router.HandleFunc(route.URI, middleware.Logger(middleware.Authenticate(route.Handle))).Methods(route.Method)
		} else {
			router.HandleFunc(route.URI, middleware.Logger(route.Handle)).Methods(route.Method)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
