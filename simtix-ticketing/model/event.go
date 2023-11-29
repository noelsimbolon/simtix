package model

import (
	"time"
)

type Event struct {
	Model
	EventName string    `json:"eventName"`
	EventTime time.Time `json:"eventTime"`
	Seats     []Seat    `json:"seats,omitempty"`
}
