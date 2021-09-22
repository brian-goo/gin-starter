package main

import (
	mdw "gtest/middleware"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	i := &lumberjack.Logger{
		Filename:   "C:/Users/81905/Desktop/gin-starter/_log/go.log",
		MaxSize:    1, // megabytes
		MaxBackups: 60,
		MaxAge:     60, //days
		Compress:   false,
	}
	// i.Rotate()
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(i)

	r := gin.Default()
	r.Use(mdw.CORS())
	// r.Use(mdw.AuthAPIKey())
	// r.Use(mdw.Auth0())

	router(r)

	server := &http.Server{
		Addr: os.Getenv("PORT"),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			if req.Method == "HEAD" {
				req.Method = "GET"
			}
			r.ServeHTTP(w, req)
		}),
		// Handler:      r,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	log.Fatal("failed to start server:", server.ListenAndServe())
}
