package controllers

import (
	"github.com/rs/zerolog/log"
	"github.com/unict-arslab/SerialController/server/model"

	"github.com/gin-gonic/gin"
)

type ServoHandlers struct {
	Client *ServoClient
}

func (h *ServoHandlers) GetAll(c *gin.Context) {
	servos, err := h.Client.GetServos()

	if err != nil {
		log.Print(err)
		c.JSON(500, "Internal Server Error")
		return
	}

	c.JSON(200, gin.H{
		"servos": servos,
	})
}

func (h *ServoHandlers) Save(c *gin.Context) {
	var servo model.Servo
	if err := c.Bind(&servo); err != nil {
		c.JSON(400, "binding error")
		return
	}

	servo.ID = ""

	created, err := h.Client.SaveTodo(servo)
	if err != nil {
		log.Print(err)
		c.JSON(500, "problem decoding body")
		return
	}

	c.JSON(201, created)
}
