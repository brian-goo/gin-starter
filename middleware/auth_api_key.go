package middleware

import (
	"gtest/handler"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AuthAPIKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token != "Bearer "+os.Getenv("API_KEY") {
			handler.LogErr(c, "auth err: wrong api key")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "not authorized"})
			return
		}

		c.Next()
	}
}
