package handler

import (
	"net/http"

	"main/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}

type User model.Task

func (h *UserHandler) GetMe(c *gin.Context) {
	var tasks []Task
	err := h.db.Find(&tasks).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}
