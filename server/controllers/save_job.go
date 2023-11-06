package controllers

import (
	"crypto/rand"
	"fmt"
	"io"

	"github.com/unict-arslab/SerialController/server/model"
)

type SaveServoJob struct {
	toSave   model.Servo
	saved    chan model.Servo
	exitChan chan error
}

func NewSaveServoJob(servo model.Servo) *SaveServoJob {
	return &SaveServoJob{
		toSave:   servo,
		saved:    make(chan model.Servo, 1),
		exitChan: make(chan error, 1),
	}
}

func (j SaveServoJob) ExitChan() chan error {
	return j.exitChan
}

func (j SaveServoJob) Run(servos map[string]model.Servo) (map[string]model.Servo, error) {
	var servo model.Servo

	id, err := newUUID()
	if err != nil {
		return nil, err
	}

	servo = model.Servo{
		ID:        id,
		Name:      j.toSave.Name,
		Min:       j.toSave.Min,
		Max:       j.toSave.Max,
		ServoType: j.toSave.ServoType,
		Robot:     j.toSave.Robot,
	}

	servos[servo.ID] = servo

	j.saved <- servo
	return servos, nil
}

// Generate a uuid to use as a unique identifier for each Todo
// http://play.golang.org/p/4FkNSiUDMg
func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}

	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80

	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
