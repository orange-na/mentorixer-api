package handler

import (
	"net/http"

	"main/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskHandler struct {
	db *gorm.DB
}

func NewTaskHandler(db *gorm.DB) *TaskHandler {
	return &TaskHandler{db: db}
}

type Task model.Task

func (h *TaskHandler) GetTasks(c *gin.Context) {
	var tasks []Task
	err := h.db.Find(&tasks).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) AddTask(c *gin.Context) {
	var task model.Task
	err := c.ShouldBindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.ID = uuid.New().String()

	err = h.db.Create(&task).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) EditTask(c *gin.Context) {
	taskID := c.Param("id")

	var task model.Task
	err := c.ShouldBindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.ID = taskID

	err = h.db.Save(&task).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
    taskID := c.Param("id")

    err := h.db.Where("id = ?", taskID).Delete(&model.Task{}).Error
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.Status(http.StatusOK)
}