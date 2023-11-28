package seat

import "github.com/shopspring/decimal"

type CreateSeatDto struct {
	EventID    string          `json:"eventID" binding:"required"`
	SeatRow    string          `json:"seatRow" binding:"required"`
	SeatNumber int             `json:"seatNumber" binding:"required,numeric"`
	Price      decimal.Decimal `json:"price"`
}

type UpdateSeatStatusDto struct {
	SeatID string     `json:"seatID" binding:"required"`
	Status SeatStatus `json:"status" binding:"required"`
}

type BookSeatDto struct {
	SeatID    string `json:"seatID" binding:"required"`
	BookingID string `json:"bookingID" binding:"required"`
}
