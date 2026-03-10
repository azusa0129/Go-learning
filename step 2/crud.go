package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Posts struct {
	Id      int    `gorm:"primaryKey" json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

var db *gorm.DB

func init() {
	var err error

	dsn := ":@tcp(localhost:3306)/asdf"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Mysql 연결 실패", err)
	}

	db.AutoMigrate(&Posts{})
}

func createPost(c *gin.Context) {
	var request PostsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, NewErrorResponse(err.Error()))
		return
	}
	post := Posts{
		Title:   request.Title,
		Content: request.Content,
		Author:  request.Author,
	}
	db.Create(&post)
	c.JSON(201, NewMessageResponse("생성되었습니다."))
}

func getPosts(c *gin.Context) {
	var posts []Posts
	db.Find(&posts)
	c.JSON(200, posts)
}

func getPost(c *gin.Context) {
	id := c.Param("id")
	var post Posts
	if err := db.First(&post, id).Error; err != nil {
		c.JSON(404, NewErrorResponse("게시물을 찾을 수 없습니다."))
		return
	}
	c.JSON(200, post)
}

func updatePost(c *gin.Context) {
	id := c.Param("id")
	var post Posts
	if err := db.First(&post, id).Error; err != nil {
		c.JSON(404, NewErrorResponse("게시물을 찾을 수 없습니다."))
		return
	}
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, NewErrorResponse(err.Error()))
		return
	}
	db.Model(&post).Updates(post)
	c.JSON(200, post)
}

func deletePost(c *gin.Context) {
	id := c.Param("id")
	var post Posts
	if err := db.Delete(&post, id).Error; err != nil {
		c.JSON(404, NewErrorResponse("게시물을 찾을 수 없습니다."))
		return
	}
	c.JSON(200, NewMessageResponse("삭제되었습니다."))
}

func main() {
	r := gin.Default()

	api := r.Group("/posts")
	{
		api.POST("", createPost)       // 게시물 생성
		api.GET("", getPosts)          // 모든 게시물 조회
		api.GET("/:id", getPost)       // 특정 게시물 조회
		api.PUT("/:id", updatePost)    // 게시물 수정
		api.DELETE("/:id", deletePost) // 게시물 삭제
	}

	r.Run(":8080")
}

type PostsRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

func NewMessageResponse(message string) MessageResponse {
	return MessageResponse{Message: message}
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{Error: message}
}
