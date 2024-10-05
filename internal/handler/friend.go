package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"main/internal/model"
	"main/pkg/db"
	"main/pkg/external/gimini"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FriendHandler struct {}

func NewFriendHandler() *FriendHandler {
	return &FriendHandler{}
}

func (h *FriendHandler) CreateFriend(c *gin.Context) {
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

	var createFriendInput struct {
		Name 			  string  `json:"name" binding:"required"`
		Mbti 			  string  `json:"mbti" binding:"required"`
		Age 		      string  `json:"age" binding:"required"`
		Gender 			  string  `json:"gender" binding:"required"`
		Description 	  *string `json:"description"` 
		ProfilePictureUrl *string `json:"profile_picture_url"`
	}

	err := c.ShouldBindJSON(&createFriendInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	age, err := strconv.Atoi(createFriendInput.Age)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid age value"})
        return
    }

	friend := model.Friend{
		UserID: u.ID,
		Name: createFriendInput.Name,
		Mbti: createFriendInput.Mbti,
		Age: age,
		Gender: createFriendInput.Gender,
		Description: createFriendInput.Description,
		ProfilePictureUrl: createFriendInput.ProfilePictureUrl,
	}

    err = db.DB.Transaction(func(tx *gorm.DB) error {
        err := tx.Create(&friend).Error
        if err != nil {
            return err
        }

        room := model.Room{UserID: u.ID, FriendID: friend.ID}
        err = tx.Create(&room).Error
        if err != nil {
            return err
        }

        return nil
    })

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	c.JSON(http.StatusCreated, friend)
}

func (h *FriendHandler) EditFriend(c *gin.Context) {
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

	err := db.DB.Where("id = ?", friendID).First(&model.Friend{}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "friend not found"})
		return
	}

    var editFriendInput struct {
        Name              string   `json:"name"`
        Mbti              string   `json:"mbti"`
        Age               string   `json:"age"`
        Gender            string   `json:"gender"`
        Description       *string  `json:"description"`
        ProfilePictureUrl *string  `json:"profile_picture_url"`
    }

	err = c.ShouldBindJSON(&editFriendInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	age, err := strconv.Atoi(editFriendInput.Age)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid age value"})
        return
    }

	friend := model.Friend{
		UserID: u.ID,
		Name: editFriendInput.Name,
		Mbti: editFriendInput.Mbti,
		Age: age,
		Gender: editFriendInput.Gender,
		Description: editFriendInput.Description,
		ProfilePictureUrl: editFriendInput.ProfilePictureUrl,
	}

	err = db.DB.Model(&model.Friend{}).Where("id = ? AND user_id = ?", friendID, u.ID).Updates(&friend).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *FriendHandler) DeleteFriend(c *gin.Context) {
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

	result := db.DB.Where("id = ? AND user_id = ?", friendID, u.ID).Delete(&model.Friend{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "friend not found"})
		return
	}

	c.Status(http.StatusOK)
}


func (h *FriendHandler) GetMessages(c *gin.Context) {
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

    var room model.Room
    err := db.DB.Where("user_id = ? AND friend_id = ?", u.ID, friendID).
        Preload("Messages").First(&room).Error
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, room.Messages)
}

func (h *FriendHandler) SendMessage(c *gin.Context) {
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

	var room model.Room
	err := db.DB.Where("user_id = ? AND friend_id = ?", u.ID, friendID).First(&room).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var sendMessageInput struct {
		Content string `json:"content" binding:"required"`
	}
	err = c.ShouldBindJSON(&sendMessageInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message := model.Message{
		RoomID: room.ID,
		UserID: &u.ID,
		Content: sendMessageInput.Content,
	}

	err = db.DB.Create(&message).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, message)
}

func (h *FriendHandler) SendMessageToGimini(c *gin.Context) {
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
	err := db.DB.Where("id = ?", friendID).First(&friend).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var room model.Room
	err = db.DB.Where("user_id = ? AND friend_id = ?", u.ID, friendID).First(&room).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var sendMessageInput struct {
		Content string `json:"content" binding:"required"`
	}
	err = c.ShouldBindJSON(&sendMessageInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var description string
	if friend.Description != nil {
		description = *friend.Description
	}

	prompt := fmt.Sprintf("以下の質問には、[%s]、[%d] 歳の [%s] 型で [%s] です。彼/彼女の特徴は次の通り: %s。その人になりきって、質問に回答してください。性格や価値観をその人の特徴に基づいて、できるだけ自然で一貫性のある応答をお願いします。」\n\n%s", friend.Name, friend.Age, friend.Mbti, friend.Gender, description, sendMessageInput.Content)
	content, err := gimini.GenerateContent(prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	message := model.Message{
		RoomID: room.ID,
		FriendID: &room.FriendID,
		Content: content,
	}

	err = db.DB.Create(&message).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, content)
}

