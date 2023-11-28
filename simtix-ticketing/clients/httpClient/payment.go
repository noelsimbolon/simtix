package httpClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"io"
	"net/http"
	"simtix-ticketing/config"
)

type PaymentClient struct {
	baseEndpoint string
}

func NewPaymentClient(config *config.Config) *PaymentClient {
	return &PaymentClient{
		baseEndpoint: config.PaymentEndpoint,
	}
}

type PostInvoicePayload struct {
	BookingID string          `json:"bookingID"`
	Amount    decimal.Decimal `json:"amount"`
}

type PostInvoiceResponse struct {
	ID         string `json:"id"`
	PaymentUrl string `json:"paymentUrl"`
}

func (c *PaymentClient) PostInvoice(payload *PostInvoicePayload) (*PostInvoiceResponse, error) {
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal invoice to JSON: %v", err)
	}

	url := fmt.Sprintf("%s/invoice", c.baseEndpoint)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadJSON))
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected response status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("unexpected response")
	}
	var data PostInvoiceResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal payload")
	}

	return &data, nil
}
