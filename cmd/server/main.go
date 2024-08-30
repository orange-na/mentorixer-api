package server

import (
	"main/internal/handler"
	"main/internal/router"
	"main/middleware"
	"main/pkg/db"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	db := db.Init()

	r.Use(middleware.SetupCORS())

	taskHandler := handler.NewTaskHandler(db)
	userHandler := handler.NewUserHandler(db)
	authHandler := handler.NewAuthHandler(db)

	router.TaskRoutes(r, taskHandler)
	router.UserRoutes(r, userHandler)
	router.AuthRoutes(r, authHandler)

	r.Run(":8080")
}