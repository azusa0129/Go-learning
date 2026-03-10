package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api := r.Group("/api/v1")
	api.POST("/hello", func(c *gin.Context) {
		var req NameRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "BadRequest",
				"message": "단어에 문제가 있습니다.",
			})
			return
		}

		hello := HelloResponse{
			Name:    req.Name,
			Message: "반갑습니다.",
		}

		c.JSON(http.StatusCreated, hello)
	})

	r.Run()
}

type NameRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}
