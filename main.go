package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Model User
type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users []User

func main() {
	r := gin.Default()

	// Route untuk mendapatkan semua user
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})

	// Route untuk membuat user baru
	r.POST("/users", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user.ID = uint(len(users) + 1)
		users = append(users, user)
		c.JSON(http.StatusCreated, user)
	})

	r.Run(":8080") // Jalankan server pada port 8080
}
