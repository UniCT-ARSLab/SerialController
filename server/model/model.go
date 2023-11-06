package model

type Servo struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Min       int64  `json:"min"`
	Max       int64  `json:"max"`
	ServoType string `json:"servoType"`
	Robot     string `json:"robot"`
}
