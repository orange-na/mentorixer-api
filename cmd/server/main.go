package server

import (
	"main/internal/handler"
	"main/internal/repository"
	"main/pkg/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	db, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// リポジトリの初期化
	taskRepo := repository.NewTaskRepository(db)
	userRepo := repository.NewUserRepository(db)

	// ハンドラーの初期化
	taskHandler := handler.NewTaskHandler(taskRepo)
	userHandler := handler.NewUserHandler(userRepo)

	// CORSの設定
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	r.Use(cors.New(config))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/tasks", func(c *gin.Context) {
		taskHandler.GetTasks(c.Writer, c.Request)
	})

	r.POST("/tasks", func(c *gin.Context) {
		taskHandler.AddTask(c.Writer, c.Request)
	})

	r.PUT("/tasks/:id", func(c *gin.Context) {
		taskHandler.EditTask(c.Writer, c.Request)
	})

	r.DELETE("/tasks/:id", func(c *gin.Context) {
		taskHandler.DeleteTask(c.Writer, c.Request)
	})

	r.GET("/users", func(c *gin.Context) {
		userHandler.GetUsers(c.Writer, c.Request)
	})

	r.POST("/sign-up", func(c *gin.Context) {
		userHandler.SignUp(c.Writer, c.Request)
	})

	r.POST("/sign-in", func(c *gin.Context) {
		userHandler.SignIn(c.Writer, c.Request)
	})

	r.Run(":8080")

}