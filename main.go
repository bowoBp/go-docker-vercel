package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// Gunakan mode release agar lebih optimal
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// Log startup
	log.Println("Starting Go server on Vercel...")

	// Endpoint utama agar tidak 404
	r.GET("/", func(c *gin.Context) {
		log.Println("Received request at /")
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to Go API on Vercel!"})
	})

	// Endpoint health check
	r.GET("/ping", func(c *gin.Context) {
		log.Println("Received request at /ping")
		c.JSON(http.StatusOK, gin.H{"message": "Pong"})
	})

	fmt.Println("Server is running at http://localhost:8080")
	r.Run(":8080")
}
