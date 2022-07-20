package routes

import "net/http"

var userRoutes = []Route{
	{
		URI:    "/users",
		Method: http.MethodPost,
		Handler: func(w http.ResponseWriter, r *http.Request) {
		},
		NeedAuth: false,
	},
	{
		URI:    "/users",
		Method: http.MethodGet,
		Handler: func(w http.ResponseWriter, r *http.Request) {
		},
		NeedAuth: false,
	},
	{
		URI:    "/users/{id}",
		Method: http.MethodGet,
		Handler: func(w http.ResponseWriter, r *http.Request) {
		},
		NeedAuth: false,
	},
	{
		URI:    "/users/{id}",
		Method: http.MethodPut,
		Handler: func(w http.ResponseWriter, r *http.Request) {
		},
		NeedAuth: false,
	},
	{
		URI:    "/users/{id}",
		Method: http.MethodDelete,
		Handler: func(w http.ResponseWriter, r *http.Request) {
		},
		NeedAuth: false,
	},
}
