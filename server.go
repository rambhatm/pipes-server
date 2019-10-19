package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/SetPipe", func(c *gin.Context) {

		user := c.PostForm("user")
		node := c.PostForm("node")
		data := c.PostForm("data")

		log.Println(user)

		SetPipe(user, node, data)
		c.JSON(http.StatusCreated, gin.H{})

	})

	router.GET("/GetPipe", func(c *gin.Context) {

		user := c.Query("user")
		node := c.Query("node")

		myPipes := GetPipe(user, node)
		if len(myPipes) == 0 {
			c.JSON(http.StatusNotFound, gin.H{})
			return
		}
		c.JSON(http.StatusFound, gin.H{"pipes": myPipes})
	})

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
}
