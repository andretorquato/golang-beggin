package routes

import (
	"api/src/controllers"
	"net/http"
)

var loginRoutes = []Route{
	{
		URI:      "/login",
		Method:   http.MethodPost,
		Handler:  controllers.Login,
		NeedAuth: false,
	},
}
