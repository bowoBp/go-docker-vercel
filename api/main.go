package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Inisialisasi router GIN
	router := gin.Default()

	// Endpoint utama
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the API"})
	})

	// Endpoint CRUD Users
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	router.POST("/users", createUser)
	router.PUT("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)

	// Gunakan PORT dari environment variable (untuk Vercel)
	port := "3000"
	router.Run(":" + port) // Gunakan port 3000 untuk Vercel
}

// Model User
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: "1", Name: "John Doe"},
	{ID: "2", Name: "Jane Doe"},
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")
	for _, u := range users {
		if u.ID == id {
			c.JSON(http.StatusOK, u)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func createUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users = append(users, newUser)
	c.JSON(http.StatusCreated, newUser)
}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	for i, u := range users {
		if u.ID == id {
			c.BindJSON(&users[i])
			c.JSON(http.StatusOK, users[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
