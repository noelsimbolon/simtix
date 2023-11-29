package dto

import "github.com/shopspring/decimal"

type CreateSeatDto struct {
	EventID    string          `json:"eventID" binding:"required"`
	SeatRow    string          `json:"seatRow" binding:"required"`
	SeatNumber int             `json:"seatNumber" binding:"required,numeric"`
	Price      decimal.Decimal `json:"price"`
}

type InvoiceStatus string

const (
	INVOICESTATUS_PAID   InvoiceStatus = "PAID"
	INVOICESTATUS_FAILED InvoiceStatus = "FAILED"
)

type UpdateSeatStatusDto struct {
	BookingID     string        `json:"bookingID" binding:"required"`
	InvoiceStatus InvoiceStatus `json:"status" binding:"required"`
}

type BookSeatDto struct {
	SeatID    string `json:"seatID" binding:"required"`
	BookingID string `json:"bookingID" binding:"required"`
}
