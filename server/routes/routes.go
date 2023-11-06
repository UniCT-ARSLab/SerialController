package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/unict-arslab/SerialController/server/controllers"
)

func Routes(router *gin.Engine, handlers *controllers.ServoHandlers) {
	router.GET("/get", handlers.GetAll)
	router.POST("/save", handlers.Save)
}
