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
		NeedAuth: true,
	},
	{
		URI:      "/users/{id}",
		Method:   http.MethodGet,
		Handler:  controllers.GetAnUser,
		NeedAuth: true,
	},
	{
		URI:      "/users/{id}",
		Method:   http.MethodPut,
		Handler:  controllers.UpdateUserData,
		NeedAuth: true,
	},
	{
		URI:      "/users/{id}",
		Method:   http.MethodDelete,
		Handler:  controllers.DeleteUser,
		NeedAuth: true,
	},
	{
		URI:      "/users/{id}/follow",
		Method:   http.MethodPost,
		Handler:  controllers.FollowUser,
		NeedAuth: true,
	},
	{
		URI:      "/users/{id}/unfollow",
		Method:   http.MethodDelete,
		Handler:  controllers.UnFollowUser,
		NeedAuth: true,
	},
	{
		URI:      "/users/{id}/followers",
		Method:   http.MethodGet,
		Handler:  controllers.GetFollowers,
		NeedAuth: true,
	},
	{
		URI:      "/users/{id}/following",
		Method:   http.MethodGet,
		Handler:  controllers.GetFollowing,
		NeedAuth: true,
	},
	{
		URI:      "/users/{id}/update-password",
		Method:   http.MethodPost,
		Handler:  controllers.UpdatePassword,
		NeedAuth: true,
	},
}
