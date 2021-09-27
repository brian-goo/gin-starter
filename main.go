package main

import (
	"gtest/endless"
	mdw "gtest/middleware"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	if os.Getenv("APP_ENV") == "DEPLOYED" {
		i := getLogger()
		// i.Rotate()
		gin.DisableConsoleColor()
		gin.DefaultWriter = io.MultiWriter(i)
		log.SetOutput(i)
	}

	r := gin.Default()
	r.Use(mdw.CORS())
	// r.Use(mdw.AuthAPIKey())
	// r.Use(mdw.Auth0())

	router(r)

	if os.Getenv("APP_ENV") != "DEPLOYED" {
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
		log.Fatal("failed to start server: ", server.ListenAndServe())
	} else {
		endless.DefaultHammerTime = -1
		endless.DefaultReadTimeOut = 60 * time.Second
		endless.DefaultWriteTimeOut = 60 * time.Second
		endless.DefaultMaxHeaderBytes = 1 << 20
		server := endless.NewServer(
			os.Getenv("PORT"),
			http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				if req.Method == "HEAD" {
					req.Method = "GET"
				}
				r.ServeHTTP(w, req)
			}),
		)
		server.ListenAndServe()
	}
}
