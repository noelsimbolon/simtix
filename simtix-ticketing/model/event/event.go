package event

import (
	"simtix-ticketing/model"
	"simtix-ticketing/model/seat"
)

type Event struct {
	model.Model
	EventName string `json:"eventName"`
	Seats []seat.Seat `json:"seats"`
}
