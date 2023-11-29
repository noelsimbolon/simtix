package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
	"simtix/clients/ticketing"
	"simtix/lib"
	"simtix/models"
	"simtix/utils/logger"
	"simtix/worker/tasks"
)

var paymentTaskCount = 0

type MakePaymentHandler struct {
	repository      *gorm.DB
	ticketingClient *ticketing.TicketingClient
}

func NewMakePaymentHandler(db *lib.Database, client *ticketing.TicketingClient) *MakePaymentHandler {
	return &MakePaymentHandler{
		repository:      db.DB,
		ticketingClient: client,
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

		paymentTaskCount += 1

		// 20% failure simulation
		if paymentTaskCount%5 == 1 {
			logger.Log.Error("Payment processing failed")
			return errors.New("Payment processing failed")
		}

		err = h.UpdateInvoiceStatus(payload.InvoiceID, models.INVOICESTATUS_PAID)
		if err != nil {
			logger.Log.Error(err)
		}

		logger.Log.Info("Payment processing successful! 🥳")
		return err
	}
}

func (h *MakePaymentHandler) UpdateInvoiceStatus(invoiceID string, status models.InvoiceStatus) error {
	var invoice models.Invoice
	log.Print(h)
	log.Print(h.repository)
	err := h.repository.Where("id = ?", invoiceID).First(&invoice).Error
	if err != nil {
		return err
	}
	invoice.Status = status

	err = h.repository.Save(invoice).Error
	if err != nil {
		return err
	}
	err = h.ticketingClient.PutBooking(&invoice)
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

func (h *MakePaymentHandler) HandleError(t *asynq.Task) error {
	payload, _ := h.unmarshalPayload(t)

	logger.Log.Info(fmt.Sprintf("Handling failed payment for invoice: %s", payload.InvoiceID))

	err := h.UpdateInvoiceStatus(payload.InvoiceID, models.INVOICESTATUS_FAILED)
	return err
}
