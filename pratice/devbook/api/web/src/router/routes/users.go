package routes

import (
	"net/http"
	"web/src/controllers"
)

var usersRoutes = []Route{
	{
		URI:      "/register",
		Method:   http.MethodGet,
		Handle:   controllers.LoadRegisterScreen,
		NeedAuth: false,
	},
	{
		URI:      "/users",
		Method:   http.MethodPost,
		Handle:   controllers.CreateUser,
		NeedAuth: false,
	},
	{
		URI:      "/users/{userId}",
		Method:   http.MethodPost,
		Handle:   controllers.CreateUser,
		NeedAuth: false,
	},
}
