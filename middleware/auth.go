package middleware

import (
	"main/internal/model"
	"main/pkg/db"
	auth "main/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)


func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			unauthorized(c, "Missing Authorization header")
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			unauthorized(c, "Invalid token format")
			return
		}

		claims, err := auth.ParseToken(token)
		if err != nil {
			unauthorized(c, "Invalid token")
			return
		}

		var user model.User
		err = db.DB.Where("id = ?", claims.UserID).First(&user).Error
		if err != nil {
			unauthorized(c, "User not found")
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func unauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": message})
	c.Abort()
}