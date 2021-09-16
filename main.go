package main

import (
	mdw "gtest/middleware"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(mdw.CORS())
	r.Use(mdw.AuthAPIKey())

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
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		// MaxHeaderBytes: ,
	}

	// if err?
	server.ListenAndServe()
}
