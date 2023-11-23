package tasks

import (
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/shopspring/decimal"
	"simtix/models"
)

const (
	TypeMakePaymentTask = "payment:make_payment"
)

var paymentTaskCount = 0

type MakePaymentPayload struct {
	InvoiceID string
	BookingID string
	Amount    decimal.Decimal
}

func NewMakePaymentTask(invoice models.Invoice) (*asynq.Task, error) {
	payload, err := json.Marshal(
		MakePaymentPayload{
			InvoiceID: invoice.ID, BookingID: invoice.BookingID,
			Amount: invoice.Amount,
		},
	)
	if err != nil {
		return nil, err
	}
	// make always fail for simulation
	return asynq.NewTask(TypeMakePaymentTask, payload, asynq.MaxRetry(0)), nil
}
