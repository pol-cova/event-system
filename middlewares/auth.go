package middlewares

import (
	"event-system/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware(context *gin.Context) {
	// Auth middleware logic
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "token required"})
		return
	}

	// Validate token logic
	userId, err := utils.ValidateToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
		return
	}

	context.Set("userId", userId)

	context.Next()

}
