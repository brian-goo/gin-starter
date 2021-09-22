package handler

import (
	"github.com/gin-gonic/gin"
)

func GetPing(c *gin.Context) {
	sub, _ := getSubFromAuth0Token(c)

	LogErr(c, "tt")

	setResponse(c, gin.H{"status": sub})
	// c.JSON(getResponse(gin.H{"status": sub}))
}
