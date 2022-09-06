package router

import "github.com/gorilla/mux"

func Generate() *mux.Router {
	return mux.NewRouter()
}
