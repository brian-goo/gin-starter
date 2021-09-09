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
	// r.Use(cors.Default())

	router(r)

	server := &http.Server{
		Addr:         os.Getenv("PORT"),
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		// MaxHeaderBytes: ,
	}

	// if err?
	server.ListenAndServe()
}
