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
		FriendRGroup.PUT("/:friend_id", friendHandler.EditFriend)
		FriendRGroup.DELETE("/:friend_id", friendHandler.DeleteFriend)
		FriendRGroup.GET("/:friend_id/messages", friendHandler.GetMessages)
		FriendRGroup.POST("/:friend_id/messages", friendHandler.SendMessage)
	}
}