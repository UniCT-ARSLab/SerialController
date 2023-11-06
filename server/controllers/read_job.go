package controllers

import (
	"github.com/unict-arslab/SerialController/server/model"
)

// Job to read all servos from the database
type ReadServosJobs struct {
	servos chan map[string]model.Servo
	exitChan chan error
}

func NewReadServosJob() *ReadServosJobs {
	return &ReadServosJobs{
		servos: make(chan map[string]model.Servo, 1),
		exitChan: make(chan error, 1),
	}
}
func (j ReadServosJobs) ExitChan() chan error {
	return j.exitChan
}

func (j ReadServosJobs) Run(servos map[string]model.Servo) (map[string]model.Servo, error) {
	j.servos <- servos

	return nil, nil
}