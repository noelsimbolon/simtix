package event

import (
	"simtix-ticketing/model"
	"simtix-ticketing/model/seat"
	"time"
)

type Event struct {
	model.Model
	EventName string      `json:"eventName"`
	EventTime time.Time   `json:"eventTime"`
	Seats     []seat.Seat `json:"seats"`
}
