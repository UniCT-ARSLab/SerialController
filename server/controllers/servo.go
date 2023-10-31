package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/unict-arslab/SerialController/server/configs"
	"github.com/unict-arslab/SerialController/server/dto"
	"github.com/unict-arslab/SerialController/server/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var servoCollection *mongo.Collection = configs.GetCollection(configs.DB, "Servos")

func SaveServo() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var servo model.Servo
		defer cancel()

		if err := c.BindJSON(&servo); err != nil {
			c.JSON(http.StatusBadRequest, dto.BadResponseDTO{Message: "Error on binding"})
			return
		}

		newServo := model.Servo{
			ID:        primitive.NewObjectID(),
			Name:      servo.Name,
			Min:       servo.Min,
			Max:       servo.Max,
			ServoType: servo.ServoType,
			Robot:     servo.Robot,
		}

		_, err := servoCollection.InsertOne(ctx, newServo)

		if err != nil {
			fmt.Println("An error occurred:", err)
			return
		}

		c.JSON(http.StatusCreated, dto.ServoResponseDTO{ServoID: newServo.ID})
	}
}