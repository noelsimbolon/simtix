package dao

import (
	"simtix-ticketing/clients/httpClient"
	"simtix-ticketing/model"
)

type BookSeatDao struct {
	Seat    model.Seat                      `json:"seat"`
	Invoice *httpClient.PostInvoiceResponse `json:"invoice"`
}
