package routes

import (
	"net/http"
	"web/src/controllers"
)

var homeRoutes = []Route{
	{
		URI:      "/home",
		Method:   http.MethodGet,
		Handle:   controllers.LoadHomeScreen,
		NeedAuth: true,
	},
}
