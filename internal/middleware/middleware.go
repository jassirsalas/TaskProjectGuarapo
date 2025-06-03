package middleware

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"strings"
	"taskproject/internal/models"

	"github.com/gin-gonic/gin"
)

var sessions = make(map[string]string)

func LoginMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var LoginUser models.LoginRequest
		if err := c.BindJSON(&LoginUser); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "username required"})
			return
		}

		token, _ := randomHex(20)
		sessions[token] = LoginUser.Username

		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		bearerToken := c.GetHeader("Authorization")
		if bearerToken == "" || !strings.HasPrefix(bearerToken, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid token"})
			return
		}

		parts := strings.SplitN(bearerToken, " ", 2)
		if len(parts) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
			return
		}

		token := parts[1]
		username, ok := sessions[token]
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		c.Set("username", username)

		c.Next()
	}
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func CreateSession(username string) (string, error) {
	token, err := randomHex(20)
	if err != nil {
		return "", err
	}
	sessions[token] = username
	return token, nil
}
