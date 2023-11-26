package seat

import "simtix-ticketing/model"

type SeatStatus string

const (
	SEATSTATUS_OPEN    SeatStatus = "OPEN"
	SEATSTATUS_ONGOING SeatStatus = "ONGOING"
	SEATSTATUS_BOOKED  SeatStatus = "BOOKED"
)

type Seat struct {
	model.Model
	EventID string     `json:"bookingID"`
	Status  SeatStatus `json:"status"`
}
