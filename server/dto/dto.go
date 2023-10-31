package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type SaveRequestoDTO struct {
	Name      string `json:"name"`
	Min       int64  `json:"min"`
	Max       int64  `json:"max"`
	ServoType string `json:"servoType"`
	Robot     string `json:"robot"`
}

type FindAllResponseDTO struct {
	Servos []ServoResponseDTO `json:"servos"`
}

type ServoResponseDTO struct {
	ServoID   primitive.ObjectID `json:"servoID"`
	Name      string             `json:"name"`
	Min       int64              `json:"min"`
	Max       int64              `json:"max"`
	ServoType string             `json:"servoType"`
	Robot     string             `json:"robot"`
}

type BadResponseDTO struct {
	Message     string             `json:"Message"`
}
