package ticketing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"simtix/lib"
	"simtix/models"
)

type TicketingClient struct {
	baseEndpoint string
}

func NewTicketingClient(config *lib.Config) *TicketingClient {
	return &TicketingClient{
		baseEndpoint: config.TicketingEndpoint,
	}
}

func (t *TicketingClient) PutBooking(invoice *models.Invoice) error {
	invoiceJSON, err := json.Marshal(invoice)
	if err != nil {
		return fmt.Errorf("failed to marshal invoice to JSON: %v", err)
	}

	url := fmt.Sprintf("%s/webhook", t.baseEndpoint)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(invoiceJSON))
	if err != nil {
		return fmt.Errorf("failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response status code: %d", resp.StatusCode)
	}

	return nil
}
