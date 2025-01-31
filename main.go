package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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

	// Ambil PORT dari environment variable atau gunakan default 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Gunakan 3000 karena Vercel default ke 3000
	}

	// Jalankan server
	log.Printf("Server running on port %s", port)
	log.Fatal(r.Run(":" + port))
}
