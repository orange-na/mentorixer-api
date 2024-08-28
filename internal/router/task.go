package router

import (
	"main/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, taskHandler *handler.Handler) {
	tasksGroup := r.Group("/tasks")
	{
		tasksGroup.GET("", taskHandler.GetTasks)
		tasksGroup.POST("", taskHandler.AddTask)
		tasksGroup.PUT("/:id", taskHandler.EditTask)
		tasksGroup.DELETE("/:id", taskHandler.DeleteTask)
	}
}