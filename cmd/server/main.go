package server

import (
	"main/internal/handler"
	"main/pkg/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	db, err := db.Init()
	if err != nil {
		panic(err)
	}

	// CORSの設定
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	taskHandler := handler.NewHandler(db)

	r.GET("/tasks", taskHandler.GetTasks)
	r.POST("/tasks", taskHandler.AddTask)
	r.PUT("/tasks/:id", taskHandler.EditTask)
	r.DELETE("/tasks/:id", taskHandler.DeleteTask)

	// r.GET("/users", func(c *gin.Context) {
	// 	userHandler.GetUsers(c.Writer, c.Request)
	// })

	// r.POST("/sign-up", func(c *gin.Context) {
	// 	userHandler.SignUp(c.Writer, c.Request)
	// })

	// r.POST("/sign-in", func(c *gin.Context) {
	// 	userHandler.SignIn(c.Writer, c.Request)
	// })

	r.Run(":8080")

}