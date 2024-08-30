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
	}
}