package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong pong ping",
		})
	})

	port := ":4001"  // Match the port with Jenkins deployment
	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
	log.Printf("Server is running at http://localhost%s", port)
}
