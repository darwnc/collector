package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	engine.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"pong": "hello"})
	})
	engine.Run(":8080")
}
