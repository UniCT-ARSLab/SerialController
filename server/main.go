package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })

	err := r.Run() // listen and serve on 0.0.0.0:8080

	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}
}
