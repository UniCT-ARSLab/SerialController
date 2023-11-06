package controllers

import (
	"github.com/unict-arslab/SerialController/server/model"
)

// client for submitting jobs and providing a repository-like interface
type ServoClient struct {
	Jobs chan Job
}

func (c *ServoClient) GetServos() ([]model.Servo, error) {
	arr := make([]model.Servo, 0)

	servos, err := c.getServosHash()
	if err != nil {
		return arr, err
	}

	for _, value := range servos {
		arr = append(arr, value)
	}

	return arr, nil
}

func (c *ServoClient) getServosHash() (map[string]model.Servo, error) {
	job := NewReadServosJob()
	c.Jobs <- job

	if err := <-job.ExitChan(); err != nil {
		return make(map[string]model.Servo, 0), err
	}

	return <-job.servos, nil
}

func (c *ServoClient) SaveTodo(servo model.Servo) (model.Servo, error) {
	job := NewSaveServoJob(servo)
	c.Jobs <- job

	if err := <-job.ExitChan(); err != nil {
		return model.Servo{}, err
	}

	return <-job.saved, nil
}
