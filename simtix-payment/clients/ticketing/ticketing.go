package ticketing

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"simtix/lib"
	"simtix/models"
)

type TicketingClient struct {
	baseEndpoint    string
	ticketingSecret string
}

func NewTicketingClient(config *lib.Config) *TicketingClient {
	return &TicketingClient{
		baseEndpoint: config.TicketingEndpoint,
	}
}

type PutBookingPayload struct {
	BookingID     string `json:"bookingID"`
	InvoiceStatus string `json:"InvoiceStatus"`
}

func (t *TicketingClient) PutBooking(invoice *models.Invoice) error {
	invoiceJSON, err := json.Marshal(invoice)
	if err != nil {
		return fmt.Errorf("failed to marshal invoice to JSON: %v", err)
	}

	url := fmt.Sprintf("%s/seat/webhook", t.baseEndpoint)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(invoiceJSON))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %v", err)
	}

	signature, err := t.getWebhookSignature(invoiceJSON)
	if err != nil {
		return fmt.Errorf("failed to calculate HMAC: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Webhook-Signature", signature)

	// Perform the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response status code: %d", resp.StatusCode)
	}

	return nil
}

func (t *TicketingClient) getWebhookSignature(payload []byte) (string, error) {
	key := []byte(t.ticketingSecret)
	h := hmac.New(sha256.New, key)
	h.Write(payload)
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature, nil
}
