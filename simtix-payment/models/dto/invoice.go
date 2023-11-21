package dto

import "github.com/shopspring/decimal"

type CreateInvoiceDto struct {
	BookingID string          `json:"bookingID" binding:"required"`
	Amount    decimal.Decimal `json:"amount" binding:"required"`
}
