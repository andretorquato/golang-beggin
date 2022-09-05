package routes

import (
	"api/src/controllers"
	"net/http"
)

var postsRoutes = []Route{
	{
		URI:      "/posts",
		Method:   http.MethodPost,
		Handler:  controllers.CreatePost,
		NeedAuth: true,
	},
	{
		URI:      "/posts",
		Method:   http.MethodGet,
		Handler:  controllers.GetAllPosts,
		NeedAuth: true,
	},
	{
		URI:      "/posts/{id}",
		Method:   http.MethodGet,
		Handler:  controllers.GetAnPost,
		NeedAuth: true,
	}, {
		URI:      "/posts/{id}",
		Method:   http.MethodPut,
		Handler:  controllers.UpdatePost,
		NeedAuth: true,
	}, {
		URI:      "/posts/{id}",
		Method:   http.MethodDelete,
		Handler:  controllers.DeletePost,
		NeedAuth: true,
	},
	{
		URI:      "/posts/users/{id}",
		Method:   http.MethodGet,
		Handler:  controllers.GetPostsByUser,
		NeedAuth: true,
	},
	{
		URI:      "/posts/{id}/like",
		Method:   http.MethodPut,
		Handler:  controllers.LikePost,
		NeedAuth: true,
	},
	{
		URI:      "/posts/{id}/dislike",
		Method:   http.MethodPut,
		Handler:  controllers.DislikePost,
		NeedAuth: true,
	},
}
