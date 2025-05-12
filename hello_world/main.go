package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Starting server on port 8080")

	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong v0.0.1",
		})
	})
	err := r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "http://localhost:8080/ping")
	if err != nil {
		fmt.Printf("Error starting server: %+v\n", err)
	}

	fmt.Println("Server started (finished?) on port 8080")
}
