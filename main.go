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

	if err := r.Run(":4000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
	log.Println("Server is running at http://localhost:4000")
}

// token tele : 7801299262:AAFTUsvVxL59EzZHQfAcdLYOgb4kK5B42Fg
// id : 6894773989
