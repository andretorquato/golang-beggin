package router

import "github.com/gorilla/mux"

func Start() *mux.Router {
	return mux.NewRouter()
}
