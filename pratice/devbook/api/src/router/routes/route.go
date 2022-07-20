package routes

import "net/http"

type Route struct {
	URI      string
	Method   string
	Handler  func(w http.ResponseWriter, r *http.Request)
	NeedAuth bool
}
