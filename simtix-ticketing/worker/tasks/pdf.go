package tasks

import (
	"encoding/json"
	"github.com/hibiken/asynq"
)

const (
	TypeGeneratePdfTask = "ticketing:generate_pdf"
)

type GeneratePdfPayload struct {
}

func NewGeneratePdfTask() (*asynq.Task, error) {
	payload, err := json.Marshal(
		GeneratePdfPayload{},
	)
	if err != nil {
		return nil, err
	}
	// make always fail for simulation
	return asynq.NewTask(TypeGeneratePdfTask, payload, asynq.MaxRetry(0)), nil
}
