package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router(r)

	// r.Run(os.Getenv("PORT"))
	r.Run("localhost:7000")
}
