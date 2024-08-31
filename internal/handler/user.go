package handler

import (
	"net/http"

	"main/internal/model"
	"main/pkg/db"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetMe(c *gin.Context) {
	user, exist := c.Get("user")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetAllFriends(c *gin.Context) {
	user, exist := c.Get("user")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		return
	}

	u, ok := user.(model.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		return
	}

	var friends []model.Friend
	err := db.DB.Where("user_id = ?", u.ID).Find(&friends).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, friends)
}

func (h *UserHandler) GetFriend(c *gin.Context) {
	user, exist := c.Get("user")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		return
	}

	u, ok := user.(model.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		return
	}

	friendID := c.Param("friend_id")
	var friend model.Friend
	err := db.DB.Where("user_id = ? AND id = ?", u.ID, friendID).First(&friend).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, friend)
}