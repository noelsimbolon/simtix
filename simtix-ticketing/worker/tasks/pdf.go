package tasks

import (
	"encoding/json"
	"github.com/hibiken/asynq"
	"simtix-ticketing/model"
)

const (
	TypeGeneratePdfTask = "ticketing:generate_pdf"
)

type GeneratePdfPayload struct {
	BookingID   string     `json:"bookingID"`
	Seat        model.Seat `json:"seat"`
	ErrorReason *string    `json:"errorReason"`
}

func NewGeneratePdfTask(payload GeneratePdfPayload) (*asynq.Task, error) {
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	// make always fail for simulation
	return asynq.NewTask(TypeGeneratePdfTask, payloadJSON, asynq.MaxRetry(0)), nil
}
