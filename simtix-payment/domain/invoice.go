package domain

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"main/lib"
	"main/models"
	"main/models/dto"
	"main/utils"
	"net/http"
)

type InvoiceService interface {
	CreateInvoice(invoiceDto dto.CreateInvoiceDto) (*models.Invoice, *utils.BaseError)
	UpdateInvoiceStatus(invoiceID string, status models.InvoiceStatus) (*models.Invoice, *utils.BaseError)
	GetInvoiceByID(invoiceID string) (
		*models.Invoice,
		*utils.BaseError,
	)
}

type InvoiceServiceImpl struct {
	repository *gorm.DB
}

func (s *InvoiceServiceImpl) CreateInvoice(invoiceDto dto.CreateInvoiceDto) (*models.Invoice, *utils.BaseError) {
	invoice := models.Invoice{
		BookingID: invoiceDto.BookingID,
		Amount:    invoiceDto.Amount,
		Status:    models.INVOICESTATUS_PENDING,
	}
	err := s.repository.Create(&invoice).Error
	if err != nil {
		return nil, &utils.BaseError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	return &invoice, nil
}

func (s *InvoiceServiceImpl) UpdateInvoiceStatus(invoiceID string, status models.InvoiceStatus) (
	*models.Invoice,
	*utils.BaseError,
) {
	invoice, baseErr := s.GetInvoiceByID(invoiceID)
	if baseErr != nil {
		return nil, baseErr
	}
	invoice.Status = status

	err := s.repository.Save(invoice).Error
	if err != nil {
		return nil, &utils.NewInternalServerError(err).BaseError
	}
	return invoice, nil
}

func (s *InvoiceServiceImpl) GetInvoiceByID(invoiceID string) (
	*models.Invoice,
	*utils.BaseError,
) {
	var invoice models.Invoice
	err := s.repository.Where("id = ? ", invoiceID).First(&invoice).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &utils.NewNotFoundError(err).BaseError
		}
		return nil, &utils.NewInternalServerError(err).BaseError
	}
	return &invoice, nil
}

func NewInvoiceService(database *lib.Database) *InvoiceServiceImpl {
	return &InvoiceServiceImpl{
		repository: database.DB,
	}
}
