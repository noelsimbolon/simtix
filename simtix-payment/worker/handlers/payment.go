package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"simtix/lib"
	"simtix/models"
	"simtix/utils/logger"
	"simtix/worker/tasks"
)

var paymentTaskCount = 0

type MakePaymentHandler struct {
	repository *gorm.DB
}

func NewMakePaymentHandler(db *lib.Database) *MakePaymentHandler {
	return &MakePaymentHandler{
		repository: db.DB,
	}
}

func (h *MakePaymentHandler) HandleMakePaymentTask() asynq.HandlerFunc {
	return func(ctx context.Context, t *asynq.Task) error {

		payload, err := h.unmarshalPayload(t)

		logger.Log.Info(
			fmt.Sprintf(
				"Processing payment for Invoice: %s with booking ID: %s and amount: %s",
				payload.InvoiceID,
				payload.BookingID,
				payload.Amount.String(),
			),
		)

		// 20% failure simulation
		if paymentTaskCount%5 == 0 {
			logger.Log.Error("Payment processing failed")
			return errors.New("Payment processing failed")
		}

		paymentTaskCount += 1

		err = h.updateInvoiceStatus(payload.InvoiceID, models.INVOICESTATUS_PAID)
		if err != nil {
			logger.Log.Error(err)
		}

		logger.Log.Info("Payment processing successful! 🥳")
		// to do : post to webhook
		return err
	}
}

func (h *MakePaymentHandler) updateInvoiceStatus(invoiceID string, status models.InvoiceStatus) error {
	var invoice models.Invoice
	err := h.repository.Where("id = ?", invoiceID).First(&invoice).Error
	if err != nil {
		return err
	}
	invoice.Status = status

	err = h.repository.Save(invoice).Error
	if err != nil {
		return err
	}
	return nil
}

func (h *MakePaymentHandler) unmarshalPayload(t *asynq.Task) (*tasks.MakePaymentPayload, error) {
	var payload tasks.MakePaymentPayload

	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		logger.Log.Error(fmt.Sprintf("json.Unrmashal failed: %v: %w", err, asynq.SkipRetry))
		return nil, err
	}
	return &payload, nil
}

func (h *MakePaymentHandler) HandleError(t *asynq.Task) {
	payload, _ := h.unmarshalPayload(t)
	logger.Log.Info(fmt.Sprintf("Handling failed payment for invoice: %s", payload.InvoiceID))
	// to do post to webhook
}
