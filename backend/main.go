package main

import (
	"fmt"
	"net/http"
	"os"

	"context"

	docs "github.com/kacarmichael/JournalBlog/backend/docs"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		return
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return
	}
	defer conn.Close(context.Background())

	fmt.Println("Hello World!")
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	r.GET("/ping", pong)

	r.GET("/users", getUser)

	r.POST("/users", createUser)

	r.PUT("/users", updateUser)

	r.DELETE("/users", deleteUser)

	r.GET("/posts", getPost)

	r.POST("/posts", createPost)

	r.PUT("/posts", updatePost)

	r.DELETE("/posts", deletePost)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run("localhost:8080")
}

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func getUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get User",
	})
}

func createUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Create User",
	})
}

func updateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Update User",
	})
}

func deleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete User",
	})
}

func getPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Post",
	})
}

func createPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Create Post",
	})
}

func updatePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Post",
	})
}

func deletePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Post",
	})
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Post struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
