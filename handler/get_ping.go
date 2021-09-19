package handler

import (
	"github.com/gin-gonic/gin"
)

func GetPing(c *gin.Context) {
	sub, _ := getSubFromAuth0Token(c)

	logErr(c, "tt")

	c.JSON(getResponse(gin.H{"status": sub}))
}
