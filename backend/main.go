package main

import (
	"fmt"
	"net/http"
	"os"

	"context"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
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

	r.GET("/ping", pong)

	r.Run("localhost:8080")
}

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
