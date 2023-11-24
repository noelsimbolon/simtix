package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type InvoiceStatus string

const (
	INVOICESTATUS_PENDING InvoiceStatus = "PENDING"
	INVOICESTATUS_PAID    InvoiceStatus = "PAID"
	INVOICESTATUS_FAILED  InvoiceStatus = "FAILED"
)

type Invoice struct {
	Model
	BookingID string          `json:"bookingID"`
	Amount    decimal.Decimal `json:"amount"`
	PaidAt    *time.Time      `json:"paidAt"`
	Status    InvoiceStatus   `json:"status"`
}
