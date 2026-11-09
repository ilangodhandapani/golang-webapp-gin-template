package main

import (
	"net/http"
	"://github.com"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
			"app":    "${{ values.name }}", // Backstage templating placeholder
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to ${{ values.name }} running on Gin!")
	})

	r.Run(":8080")
}
