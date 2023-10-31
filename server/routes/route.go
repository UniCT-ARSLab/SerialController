package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/unict-arslab/SerialController/server/controllers"
)

func Routes(router *gin.Engine) {
	router.POST("/saveServo", controllers.SaveServo())
}
