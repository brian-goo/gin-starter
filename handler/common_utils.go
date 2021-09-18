package handler

import (
	"errors"

	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
)

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
