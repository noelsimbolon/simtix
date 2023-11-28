package invoice

import (
	"fmt"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"simtix/lib"
	"simtix/models"
	"simtix/models/dao"
	"simtix/models/dto"
	"simtix/utils"
	"simtix/worker"
	"simtix/worker/tasks"
)

type InvoiceService interface {
	CreateInvoice(invoiceDto dto.CreateInvoiceDto) (*dao.CreateInvoiceDao, *utils.BaseError)
	UpdateInvoiceStatus(invoiceID string, status models.InvoiceStatus) (*models.Invoice, *utils.BaseError)
	GetInvoiceByID(invoiceID string) (
		*models.Invoice,
		*utils.BaseError,
	)
}

type InvoiceServiceImpl struct {
	config       *lib.Config
	repository   *gorm.DB
	workerClient *worker.WorkerClient
}

func NewInvoiceService(
	database *lib.Database, config *lib.Config, client *worker.WorkerClient,
) *InvoiceServiceImpl {
	return &InvoiceServiceImpl{
		repository:   database.DB,
		config:       config,
		workerClient: client,
	}
}

func (s *InvoiceServiceImpl) CreateInvoice(invoiceDto dto.CreateInvoiceDto) (
	*dao.CreateInvoiceDao,
	*utils.BaseError,
) {
	isInvoiceExist, err := s.checkExistingInvoice(invoiceDto.BookingID)
	if isInvoiceExist {
		return nil, ErrBookingExist
	}

	if err != nil {
		return nil, DbErrGetInvoice
	}

	invoice := models.Invoice{
		BookingID: invoiceDto.BookingID,
		Amount:    invoiceDto.Amount,
		Status:    models.INVOICESTATUS_PENDING,
	}

	err = s.repository.Create(&invoice).Error
	if err != nil {
		return nil, DbErrCreateInvoice
	}

	paymentUrl := s.generatePaymentUrl(invoice.ID)

	err = s.enqueuePaymentTask(invoice)
	if err != nil {
		return nil, ErrEnqueuePaymentTaskFail
	}

	return &dao.CreateInvoiceDao{
		ID:         invoice.ID,
		PaymentUrl: paymentUrl,
	}, nil
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
		return nil, DbErrUpdateInvoice
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
			return nil, ErrInvoiceNotFound
		}
		return nil, DbErrGetInvoice
	}
	return &invoice, nil
}

func (s *InvoiceServiceImpl) checkExistingInvoice(bookingID string) (bool, error) {
	var existingInvoice models.Invoice
	// check existing invoice with status paid / pending
	// if failed, can retry?
	err := s.repository.Where(
		"booking_id = ? AND status = ANY(?::_invoice_status)", bookingID,
		pq.Array([]models.InvoiceStatus{models.INVOICESTATUS_PAID, models.INVOICESTATUS_PENDING}),
	).First(&existingInvoice).Error

	// record found
	if err == nil {
		return true, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return true, err
	}

	return false, nil
}

func (s *InvoiceServiceImpl) generatePaymentUrl(invoiceID string) string {
	return fmt.Sprintf("%s:%d/payment?invoiceID=%s", s.config.ServiceHost, s.config.ServicePort, invoiceID)
}

func (s *InvoiceServiceImpl) enqueuePaymentTask(invoice models.Invoice) error {
	paymentTask, err := tasks.NewMakePaymentTask(invoice)
	if err != nil {
		return err
	}
	s.workerClient.Enqueue(paymentTask)
	return nil
}
