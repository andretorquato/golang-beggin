package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

func Start() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
