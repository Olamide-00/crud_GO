package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello there"})
	})

	router.GET("/about", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "about us"})
	})

	router.GET("/contact", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "about us"})
	})

	router.Run(":8080")
}
