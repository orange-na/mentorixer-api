package router

import (
	"main/internal/handler"
	"main/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, userHandler *handler.UserHandler) {
	UserGroup := r.Group("/users")
	UserGroup.Use(middleware.JwtAuthMiddleware())
	{
		UserGroup.GET("/me", userHandler.GetMe)
		UserGroup.GET("/me/friends", userHandler.GetAllFriends)
		UserGroup.GET("/me/friends/:friend_id", userHandler.GetFriend)
	}
}