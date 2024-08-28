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
	db, err := db.Init()
	if err != nil {
		panic(err)
	}
	
	r.Use(middleware.SetupCORS())

	taskHandler := handler.NewHandler(db)
	router.SetupRoutes(r, taskHandler)

	r.Run(":8080")
}