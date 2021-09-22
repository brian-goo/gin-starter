package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getResponse(res interface{}) (int, interface{}) {
	return http.StatusOK, res
}

// func getErrorResponse(res interface{}) (int, interface{}) {
// 	return http.StatusInternalServerError, res
// }

func setResponse(c *gin.Context, res interface{}) {
	c.JSON(getResponse(res))
}

// func setErrorResponse(c *gin.Context, res interface{}) {
// 	c.JSON(getErrorResponse(res))
// }
