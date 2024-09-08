package server

import (
	"fmt"
	"main/internal/handler"
	"main/internal/router"
	"main/middleware"
	"main/pkg/db"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Run() {
    err := godotenv.Load()
    if err != nil {
		panic(err)
    }


	r := gin.Default()
	db.Init()

	r.Use(middleware.SetupCORS())

	userHandler := handler.NewUserHandler()
	friendHandler := handler.NewFriendHandler()
	authHandler := handler.NewAuthHandler()

	router.UserRoutes(r, userHandler)
	router.FriendRoutes(r, friendHandler)
	router.AuthRoutes(r, authHandler)

	port := os.Getenv("PORT")
	r.Run(fmt.Sprintf(":%s", port))
}