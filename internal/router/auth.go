package router

import (
	"main/internal/handler"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, authHandler *handler.AuthHandler) {
	AuthGroup := r.Group("")
	{
		AuthGroup.POST("/sign-up", authHandler.SignUp)
		AuthGroup.POST("/sign-in", authHandler.SignIn)
	}
}