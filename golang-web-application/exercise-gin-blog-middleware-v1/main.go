package main

import (
	"encoding/base64"
	"strconv"
	"strings"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var Posts = []Post{
	{ID: 1, Title: "Judul Postingan Pertama", Content: "Ini adalah postingan pertama di blog ini.", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 2, Title: "Judul Postingan Kedua", Content: "Ini adalah postingan kedua di blog ini.", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []User{
	{Username: "user1", Password: "pass1"},
	{Username: "user2", Password: "pass2"},
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: answer here
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		auth := strings.SplitN(authHeader, " ", 2)
		if len(auth) != 2 || auth[0] != "Basic" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		payload, err := base64.StdEncoding.DecodeString(auth[1])
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		creds := strings.SplitN(string(payload), ":", 2)
		if len(creds) != 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		username := creds[0]
		password := creds[1]

		// Check if credentials are valid
		valid := false
		for _, user := range users {
			if user.Username == username && user.Password == password {
				valid = true
				break
			}
		}

		if !valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Jika semuanya valid, lanjutkan ke handler selanjutnya
		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//Set up authentication middleware here // TODO: replace this
	r.Use(authMiddleware())

	r.GET("/posts", func(c *gin.Context) {
		// TODO: answer here
		idStr := c.Query("id")
		if idStr != "" {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
				return
			}

			for _, post := range Posts {
				if post.ID == id {
					c.JSON(http.StatusOK, gin.H{"post": post})
					return
				}
			}

			c.JSON(http.StatusNotFound, gin.H{"error": "Postingan tidak ditemukan"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"posts": Posts})
	})

	r.POST("/posts", func(c *gin.Context) {
		// TODO: answer here
		var newPost Post
		if err := c.ShouldBindJSON(&newPost); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		newPost.ID = len(Posts) + 1
		newPost.CreatedAt = time.Now()
		newPost.UpdatedAt = time.Now()
		Posts = append(Posts, newPost)

		c.JSON(http.StatusCreated, gin.H{"message": "Postingan berhasil ditambahkan", "post": newPost})
	})
	return r
}

func main() {
	r := SetupRouter()

	r.Run(":8080")
}
