package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:      "/users",
		Method:   http.MethodPost,
		Handler:  controllers.CreateUser,
		NeedAuth: false,
	},
	{
		URI:      "/users",
		Method:   http.MethodGet,
		Handler:  controllers.GetAllUsers,
		NeedAuth: false,
	},
	{
		URI:      "/users/{id}",
		Method:   http.MethodGet,
		Handler:  controllers.GetAnUser,
		NeedAuth: false,
	},
	{
		URI:      "/users/{id}",
		Method:   http.MethodPut,
		Handler:  controllers.UpdateUserData,
		NeedAuth: false,
	},
	{
		URI:      "/users/{id}",
		Method:   http.MethodDelete,
		Handler:  controllers.DeleteUser,
		NeedAuth: false,
	},
}
