package utils

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userId := session.Get("userId")

		if userId == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authentication required!"})
			return
		}
		c.Set("userId", userId)
		c.Next()
	}
}
