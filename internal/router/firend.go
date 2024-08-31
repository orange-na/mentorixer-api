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
		FriendRGroup.POST("/", friendHandler.CreateFriend)
		FriendRGroup.PUT("/:id", friendHandler.EditFriend)
		FriendRGroup.DELETE("/:id", friendHandler.DeleteFriend)
	}
}