package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPing(c *gin.Context) {
	user := c.Request.Context().Value("user")
	log.Println(user)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
