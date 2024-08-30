package router

import (
	"main/internal/handler"
	"main/middleware"

	"github.com/gin-gonic/gin"
)

func FriendRoutes(r *gin.Engine, friendHandler *handler.FriendHandler) {
	FriendRGroup := r.Group("/friends")
	FriendRGroup.Use(middleware.JwtAuthMiddleware())
	{
		FriendRGroup.GET("/", friendHandler.GetAllFriends)
	}
}