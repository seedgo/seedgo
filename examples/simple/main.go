package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/seedgo/seedgo"
)

func main() {
	server := seedgo.NewServer()

	server.GetEngine().GET("/custom", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "custom endpoint"})
	})

	err := server.Start()
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
