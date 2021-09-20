package handler

import (
	"errors"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
)

func logErr(c *gin.Context, err interface{}) {
	switch e := err.(type) {
	// this is for standard error type eg. from db lib
	case error:
		c.Error(e)
	case string:
		c.Error(errors.New(e))
	default:
		c.Error(errors.New("error occurred: details unknown"))
	}
}

func getSubFromAuth0Token(c *gin.Context) (string, error) {
	token, ok := c.Request.Context().Value("user").(*jwt.Token)
	if !ok {
		return "", errors.New("failed to get token from request context")
	}

	if sub, ok := token.Claims.(jwt.MapClaims)["sub"].(string); ok && token.Valid {
		return sub, nil
	} else {
		return "", errors.New("failed to get sub from auth0 token")
	}
}
