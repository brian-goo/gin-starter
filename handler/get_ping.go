package handler

import (
	"log"
	"net/http"

	jwt "github.com/form3tech-oss/jwt-go"

	"github.com/gin-gonic/gin"
)

func GetPing(c *gin.Context) {
	user := c.Request.Context().Value("user")
	token, _ := user.(*jwt.Token)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Println(claims["sub"])

	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
