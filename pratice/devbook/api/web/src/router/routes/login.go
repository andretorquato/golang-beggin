package routes

import (
	"net/http"
	"web/src/controllers"
)

var loginRoutes = []Route{
	{
		URI:      "/",
		Method:   http.MethodGet,
		Handle:   controllers.LoadLoginScreen,
		NeedAuth: false,
	},
	{
		URI:      "/login",
		Method:   http.MethodGet,
		Handle:   controllers.LoadLoginScreen,
		NeedAuth: false,
	},
}
