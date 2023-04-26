package api

import (
	"playpost/core/api/controllers"
	"playpost/core/api/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			auth := new(controllers.AuthController)
			v1.POST("/auth", auth.Auth)
			v1.POST("/auth/register", auth.Register)
			v1.POST("/auth/reset_password", auth.ResetPassword)

			v1.Use(middlewares.AuthMiddleware())

			v1.POST("/auth/logout", auth.Logout)

			user := new(controllers.UserController)
			v1.GET("/user", user.GetRecommended)
			v1.GET("/user/:id", user.GetDetails)
			v1.GET("/user/search/:query", user.Search)
			v1.POST("/user/:id/friend", user.CreateFriendRequest)
			v1.PATCH("/user/:id/friend", user.RespondFriendRequest)
			v1.DELETE("/user/:id/friend", user.CancelFriendRequest)

			post := new(controllers.PostController)
			v1.POST("/post", post.CreatePost)
			v1.GET("/post/:id", post.GetPost)

			comment := new(controllers.CommentController)
			v1.POST("/post/:id/comment", comment.CommentOnPost)
			v1.POST("/comment/:id", comment.ReplyComment)
			v1.DELETE("/comment/:id", comment.DeleteComment)

			like := new(controllers.LikeController)
			v1.POST("/post/:id/like", like.LikePost)
			v1.POST("/comment/:id/like", like.LikeComment)
		}
	}
	return router
}
