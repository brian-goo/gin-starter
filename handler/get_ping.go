package handler

import (
	"github.com/gin-gonic/gin"
)

func GetPing(c *gin.Context) {
	sub, _ := getSubFromAuth0Token(c)

	c.JSON(getResponse(gin.H{"status": sub}))
}
