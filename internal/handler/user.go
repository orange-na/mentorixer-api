package handler

import (
	"net/http"

	"main/internal/model"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

type User model.User

func (h *UserHandler) GetMe(c *gin.Context) {
	user, exist := c.Get("user")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
