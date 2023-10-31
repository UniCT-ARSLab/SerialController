package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/unict-arslab/SerialController/server/configs"
	"github.com/unict-arslab/SerialController/server/routes"
)

func main() {
	router := gin.Default()

	//run database
	configs.ConnectDB()

	//routes
	routes.Routes(router)

	err := router.Run("localhost:8080")

	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}
}
