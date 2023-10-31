package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Servo struct {
	ID        primitive.ObjectID `json:"_id,omitempty"`
	Name      string             `json:"name,omitempty"`
	Min       int64              `json:"min,omitempty"`
	Max       int64              `json:"max,omitempty"`
	ServoType string             `json:"servoType,omitempty"`
	Robot     string             `json:"robot,omitempty"`
}
