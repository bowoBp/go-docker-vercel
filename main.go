package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User
var nextID uint = 1

func main() {
	r := gin.Default()

	// Endpoint untuk homepage agar tidak 404
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to Go API on Vercel!"})
	})

	// Endpoint untuk healthy check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Pong"})
	})

	// Endpoint tambahan lainnya
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})

	// Ambil PORT dari environment variable atau gunakan default 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	log.Fatal(r.Run("0.0.0.0:8080")) // Harus menggunakan 0.0.0.0 di Docker

}
