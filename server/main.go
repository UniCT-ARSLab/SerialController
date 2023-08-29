package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	err := r.Run() // listen and serve on 0.0.0.0:8080

	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}
}
