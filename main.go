package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: "1", Name: "John Doe"},
	{ID: "2", Name: "Jane Doe"},
}

func getUsers(c *gin.Context) {
	c.JSON(200, users)
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")
	for _, u := range users {
		if u.ID == id {
			c.JSON(200, u)
			return
		}
	}
	c.JSON(404, gin.H{"message": "User not found"})
}

func createUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	users = append(users, newUser)
	c.JSON(201, newUser)
}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	for i, u := range users {
		if u.ID == id {
			c.BindJSON(&users[i])
			c.JSON(200, users[i])
			return
		}
	}
	c.JSON(404, gin.H{"message": "User not found"})
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(200, gin.H{"message": "User deleted"})
			return
		}
	}
	c.JSON(404, gin.H{"message": "User not found"})
}

func main() {
	r := gin.Default()
	r.GET("/users", getUsers)
	r.GET("/users/:id", getUserByID)
	r.POST("/users", createUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

	r.Run() // Port default 8080
}
