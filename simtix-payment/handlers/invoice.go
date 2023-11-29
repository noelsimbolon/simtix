package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simtix/domain/invoice"
	"simtix/models/dto"
	"simtix/utils/logger"
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
		logger.Log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	invoice, customError := h.service.CreateInvoice(invoiceDto)
	if customError != nil {
		logger.Log.Error(customError.Err.Error())
		c.AbortWithStatusJSON(customError.StatusCode, gin.H{"error": customError.Err.Error()})
		return
	}

	c.JSON(http.StatusCreated, invoice)
}

func NewInvoiceHandlerImpl(service invoice.InvoiceService) *InvoiceHandlerImpl {
	return &InvoiceHandlerImpl{
		service: service,
	}
}
