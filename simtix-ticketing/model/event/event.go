package event

import "simtix-ticketing/model"

type Event struct {
	model.Model
	EventName string `json:"eventName"`
}
