package main

import (
	"gtest/handler"

	"github.com/gin-gonic/gin"
)

func router(r *gin.Engine) {
	r.GET("/ping", handler.GetPing)
	// r.HEAD("/ping", handler.GetPing)

	r.POST("/ping", handler.GetPing)
}
