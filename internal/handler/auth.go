package handler

import (
	"net/http"

	"main/internal/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	// "github.com/dgrijalva/jwt-go"
)

type AuthHandler struct {
	db *gorm.DB
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{db: db}
}

type Auth model.Task

func (h *AuthHandler) SignUp(c *gin.Context) {
	var signUpInput struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	err := c.ShouldBindJSON(&signUpInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser model.User
	err = h.db.Where("email = ?", signUpInput.Email).First(&existingUser).Error
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
		return
	}

	EncryptedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Name: signUpInput.Name,
		Email: signUpInput.Email,
		EncryptedPassword: string(EncryptedPassword),
	}

	err = h.db.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *AuthHandler) SignIn(c *gin.Context) {
}
