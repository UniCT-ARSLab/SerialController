package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/unict-arslab/SerialController/server/controllers"
	"github.com/unict-arslab/SerialController/server/routes"
)

func main() {
	db := "./db.json"

	// create channel to communicate over
	jobs := make(chan controllers.Job)

	// start watching jobs channel for work
	go controllers.ProcessJobs(jobs, db)

	// create client for submitting jobs / providing interface to db
	client := &controllers.ServoClient{Jobs: jobs}
	handlers := &controllers.ServoHandlers{Client: client}

	//routes
	router := gin.Default()

	routes.Routes(router, handlers)

	//start web server
	err := router.Run("localhost:8080")

	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}
}
