package model

import (
	"github.com/shopspring/decimal"
)

type SeatStatus string

const (
	SEATSTATUS_OPEN    SeatStatus = "OPEN"
	SEATSTATUS_ONGOING SeatStatus = "ONGOING"
	SEATSTATUS_BOOKED  SeatStatus = "BOOKED"
)

type Seat struct {
	Model
	EventID    string          `json:"eventID"`
	Status     SeatStatus      `json:"status"`
	SeatRow    string          `json:"seatRow"`
	SeatNumber int             `json:"seatNumber"`
	Price      decimal.Decimal `json:"price"`
	Event      Event           `json:"event"`
	BookingID  *string          `json:"bookingID"`
}
