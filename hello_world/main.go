package main

// Sources
// https://medium.com/geekculture/full-stack-application-with-go-gin-react-and-mongodb-37b63ef71133
// https://www.youtube.com/watch?v=bj77B59nkTQ&t=562s

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Ball struct {
	Name   string `json:"name"`
	Color  string `json:"color"`
	Amount int    `json:"amount"`
}

var balls = []Ball{
	{Name: "Football", Color: "White", Amount: 3},
	{Name: "Basket", Color: "Orange", Amount: 7},
}

func addBall(ball Ball) {
	for _, eb := range balls {
		if ball.Name == eb.Name {
			eb.Amount += ball.Amount
			return
		}
	}
	balls = append(balls, ball)
}

func postMethod(c *gin.Context) {
	var ball Ball

	if err := c.BindJSON(&ball); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Obtained ball: %+v", ball)
	addBall(ball)
	fmt.Println("Added ball succesfully. About to respond.")
	c.JSON(http.StatusCreated, ball)
}

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
	r.GET("/inventory", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, balls)
	})
	r.POST("/throw", postMethod)
	err := r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "http://localhost:8080/ping")
	if err != nil {
		fmt.Printf("Error starting server: %+v\n", err)
	}

	fmt.Println("Server started (finished?) on port 8080")
}
