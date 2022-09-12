package routes

import (
	"net/http"

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

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Handle).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
