package main

import (
	"log"
	"os"

	"github.com/bebek-goreng/golang-jwt-auth/initializer"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnv()
	initializer.ConnectDb()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set in env file")
	}

	r := gin.Default()

	r.Run(":" + port)
}
