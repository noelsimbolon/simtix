package httpClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"io"
	"log"
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

type ErrorResponse struct {
	error string `json:"error"`
}

func (c *PaymentClient) PostInvoice(payload *PostInvoicePayload) (*PostInvoiceResponse, error) {
	payloadJSON, err := json.Marshal(payload)

	log.Print(payloadJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal invoice to JSON: %v", err)
	}

	url := fmt.Sprintf("%s/invoice", c.baseEndpoint)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadJSON))
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		var testt interface{}
		json.Unmarshal(body, testt)
		log.Printf("%v", testt)
		log.Print()
		return nil, fmt.Errorf("unexpected response")
	}
	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected response status code: %d", resp.StatusCode)
	}

	var data PostInvoiceResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal payload")
	}

	return &data, nil
}
