package router

import (
	"main/internal/handler"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, userHandler *handler.UserHandler) {
	UserGroup := r.Group("/users")
	{
		UserGroup.GET("", userHandler.GetMe)
	}
}