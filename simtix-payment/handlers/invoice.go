package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simtix/domain/invoice"
	"simtix/models/dto"
)

type InvoiceHandler interface {
	PostInvoice(c *gin.Context)
}

type InvoiceHandlerImpl struct {
	service invoice.InvoiceService
}

func (h *InvoiceHandlerImpl) PostInvoice(c *gin.Context) {
	var invoiceDto dto.CreateInvoiceDto
	err := c.ShouldBindJSON(&invoiceDto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	invoice, customError := h.service.CreateInvoice(invoiceDto)
	if customError != nil {
		c.AbortWithStatusJSON(customError.StatusCode, gin.H{"error": customError.Err})
		return
	}

	c.JSON(http.StatusCreated, invoice)
}

func NewInvoiceHandlerImpl(service invoice.InvoiceService) *InvoiceHandlerImpl {
	return &InvoiceHandlerImpl{
		service: service,
	}
}
