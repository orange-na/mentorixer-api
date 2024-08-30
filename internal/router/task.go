package router

import (
	"main/internal/handler"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(r *gin.Engine, taskHandler *handler.TaskHandler) {
	taskGroup := r.Group("/tasks")
	{
		taskGroup.GET("", taskHandler.GetTasks)
		taskGroup.POST("", taskHandler.AddTask)
		taskGroup.PUT("/:id", taskHandler.EditTask)
		taskGroup.DELETE("/:id", taskHandler.DeleteTask)
	}
}