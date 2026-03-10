package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	api := r.Group("/api/v1")
	api.GET("/posts", getPosts)

	r.Run(":8080")
}

func getPosts(c *gin.Context) {
	c.JSON(200, gin.H{
		"title":   "test",
		"content": "test123",
	})
}
