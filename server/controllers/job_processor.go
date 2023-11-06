package controllers

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/unict-arslab/SerialController/server/model"
)

type Job interface {
	Run(servos map[string]model.Servo) (map[string]model.Servo, error)
	ExitChan() chan error
}

func ProcessJobs(jobs chan Job, db string) {
	for {
		j := <-jobs

		// Read the database
		servos := make(map[string]model.Servo, 0)
		content, err := ioutil.ReadFile(db)

		if err == nil {
			if err = json.Unmarshal(content, &servos); err == nil {

				// Run the job
				servosMod, err := j.Run(servos)

				// If there were modifications, write them back to the database
				if err == nil && servosMod != nil {
					data, err := json.Marshal(servosMod)
					if err == nil {
						err = os.WriteFile(db, data, 0644)
					}
				}
			}
		}

		j.ExitChan() <- err
	}
}
