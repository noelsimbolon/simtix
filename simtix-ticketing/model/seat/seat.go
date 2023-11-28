package seat

import (
	"github.com/shopspring/decimal"
	"simtix-ticketing/model"
)

type SeatStatus string

const (
	SEATSTATUS_OPEN    SeatStatus = "OPEN"
	SEATSTATUS_ONGOING SeatStatus = "ONGOING"
	SEATSTATUS_BOOKED  SeatStatus = "BOOKED"
)

type Seat struct {
	model.Model
	EventID    string          `json:"eventID"`
	Status     SeatStatus      `json:"status"`
	SeatRow    string          `json:"seatRow"`
	SeatNumber int             `json:"seatNumber"`
	Price      decimal.Decimal `json:"price"`
}
