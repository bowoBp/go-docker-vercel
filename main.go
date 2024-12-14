package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// User representasi data user dalam aplikasi
type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Untuk menyimpan data pengguna sementara
var users []User
var nextID uint = 1

func main() {
	r := gin.Default()

	// Endpoint untuk healthy check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Pong"})
	})

	// Endpoint untuk mengambil seluruh user
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
		return
	})

	// Endpoint untuk mengambil user berdasarkan ID
	r.GET("/users/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
			return
		}

		// Mencari user berdasarkan ID
		for _, user := range users {
			if user.ID == uint(id) {
				c.JSON(http.StatusOK, user)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
	})

	// Endpoint untuk menambahkan user baru
	r.POST("/users", func(c *gin.Context) {
		var newUser User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Menetapkan ID unik untuk pengguna baru
		newUser.ID = nextID
		nextID++

		// Menambahkan pengguna ke dalam slice
		users = append(users, newUser)
		c.JSON(http.StatusCreated, newUser)
	})

	// Endpoint untuk mengupdate user berdasarkan ID
	r.PUT("/users/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
			return
		}

		var updatedUser User
		if err := c.ShouldBindJSON(&updatedUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Mencari user berdasarkan ID
		for i, user := range users {
			if user.ID == uint(id) {
				users[i].Name = updatedUser.Name
				users[i].Age = updatedUser.Age
				c.JSON(http.StatusOK, users[i])
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
	})

	// Endpoint untuk menghapus user berdasarkan ID
	r.DELETE("/users/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
			return
		}

		// Mencari dan menghapus user berdasarkan ID
		for i, user := range users {
			if user.ID == uint(id) {
				users = append(users[:i], users[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "User dihapus"})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
	})

	r.Run(":8080")
}
