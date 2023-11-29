package seat

import (
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"log"
	"simtix-ticketing/clients/httpClient"
	"simtix-ticketing/config"
	"simtix-ticketing/database"
	"simtix-ticketing/model"
	"simtix-ticketing/model/dao"
	"simtix-ticketing/model/dto"
	"simtix-ticketing/utils"
	"simtix-ticketing/worker"
	"simtix-ticketing/worker/tasks"
)

type SeatService interface {
	GetSeatsByEventID(eventID string) ([]model.Seat, *utils.Error)
	GetSeatByID(id string) (*model.Seat, *utils.Error)
	CreateSeat(dto dto.CreateSeatDto) (*model.Seat, *utils.Error)
	BookSeat(updateSeatStatusDto dto.BookSeatDto) (*dao.BookSeatDao, *utils.Error)
	UpdateSeatStatus(updateSeatStatusDto dto.UpdateSeatStatusDto) (*model.Seat, *utils.Error)
}

type SeatServiceImpl struct {
	config        *config.Config
	repository    *gorm.DB
	paymentClient *httpClient.PaymentClient
	workerClient  *worker.WorkerClient
}

func NewSeatService(
	config *config.Config, database *database.Database, client *httpClient.PaymentClient,
	workerClient *worker.WorkerClient,
) *SeatServiceImpl {
	return &SeatServiceImpl{
		config:        config,
		repository:    database.DB,
		paymentClient: client,
		workerClient:  workerClient,
	}
}

func (s *SeatServiceImpl) GetSeatsByEventID(eventID string) ([]model.Seat, *utils.Error) {
	var seats []model.Seat
	err := s.repository.Where("event_id = ?", eventID).Find(&seats).Error
	if err != nil {
		return nil, DbErrGetSeats
	}
	return seats, nil
}

func (s *SeatServiceImpl) GetSeatByID(id string) (*model.Seat, *utils.Error) {
	var seat model.Seat
	err := s.repository.Preload("Event").Where("id = ?", id).First(&seat).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrSeatNotFound
		}
		return nil, DbErrGetSeat
	}
	return &seat, nil
}

func (s *SeatServiceImpl) GetSeatByBookingID(id string) (*model.Seat, *utils.Error) {
	var seat model.Seat
	err := s.repository.Preload("Event").Where("booking_id = ?", id).First(&seat).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrSeatNotFound
		}
		return nil, DbErrGetSeat
	}
	return &seat, nil
}

func (s *SeatServiceImpl) CreateSeat(seatDto dto.CreateSeatDto) (*model.Seat, *utils.Error) {
	var event model.Event
	dbErr := s.repository.Where("id = ?", seatDto.EventID).First(&event).Error
	if dbErr != nil {
		return nil, ErrEventNotExist
	}

	var existingSeat model.Seat
	err := s.repository.Where(
		"event_id = ? AND seat_number = ? AND seat_row = ?", seatDto.EventID, seatDto.SeatNumber,
		seatDto.SeatRow,
	).First(&existingSeat).Error

	if err == nil {
		return nil, ErrSeatExists
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, DbErrGetSeat
	}

	seat := model.Seat{
		EventID:    seatDto.EventID,
		SeatNumber: seatDto.SeatNumber,
		SeatRow:    seatDto.SeatRow,
		Status:     model.SEATSTATUS_OPEN,
		Price:      seatDto.Price,
	}
	err = s.repository.Create(&seat).Error
	if err != nil {
		return nil, DbErrCreateSeat
	}

	return &seat, nil
}

// for webhook
func (s *SeatServiceImpl) UpdateSeatStatus(updateSeatStatusDto dto.UpdateSeatStatusDto) (
	*model.Seat, *utils.Error,
) {
	existingSeat, err := s.GetSeatByBookingID(updateSeatStatusDto.BookingID)
	if err != nil {
		return nil, err
	}

	if updateSeatStatusDto.InvoiceStatus == dto.INVOICESTATUS_PAID {
		existingSeat.Status = model.SEATSTATUS_BOOKED
	} else if updateSeatStatusDto.InvoiceStatus == dto.INVOICESTATUS_FAILED {
		existingSeat.Status = model.SEATSTATUS_OPEN
	}

	dbErr := s.repository.Save(&existingSeat).Error
	if dbErr != nil {
		return nil, DbErrCreateSeat
	}

	enqueueErr := s.enqueuePdfTask(existingSeat, updateSeatStatusDto.BookingID)
	if enqueueErr != nil {
		return nil, err
	}
	return existingSeat, nil
}

// for client service
func (s *SeatServiceImpl) BookSeat(updateSeatStatusDto dto.BookSeatDto) (
	*dao.BookSeatDao, *utils.Error,
) {
	existingSeat, err := s.GetSeatByID(updateSeatStatusDto.SeatID)
	if err != nil {
		return nil, err
	}
	if existingSeat.Status != model.SEATSTATUS_OPEN {
		return nil, ErrSeatNotAvailable
	}
	if !s.checkSeatExternally(updateSeatStatusDto.SeatID) {
		return nil, ErrExternalCallFailed
	}
	invoice, err := s.makeInvoice(existingSeat.Price, updateSeatStatusDto.BookingID)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	// what to do with invoice?
	existingSeat.Status = model.SEATSTATUS_ONGOING
	existingSeat.BookingID = &updateSeatStatusDto.BookingID
	dbErr := s.repository.Save(&existingSeat).Error
	if dbErr != nil {
		log.Print(dbErr)
		return nil, DbErrCreateSeat
	}

	bookSeatDao := dao.BookSeatDao{
		Seat:    *existingSeat,
		Invoice: invoice,
	}
	return &bookSeatDao, nil
}

var checkSeatExternallyCnt = 0

// 20% failure rate simulation
func (s *SeatServiceImpl) checkSeatExternally(seatID string) bool {
	checkSeatExternallyCnt += 1
	return checkSeatExternallyCnt%5 != 0
}

func (s *SeatServiceImpl) makeInvoice(amount decimal.Decimal, bookingID string) (
	*httpClient.PostInvoiceResponse, *utils.Error,
) {
	payload := httpClient.PostInvoicePayload{
		BookingID: bookingID,
		Amount:    amount,
	}
	invoice, err := s.paymentClient.PostInvoice(&payload)
	if err != nil {
		log.Print(err)
		return nil, ErrMakePaymentFailed
	}
	return invoice, nil
}

func (s *SeatServiceImpl) enqueuePdfTask(seat *model.Seat, bookingID string) error {
	payload := tasks.GeneratePdfPayload{
		BookingID: bookingID,
		Seat:      *seat,
	}
	pdfTask, err := tasks.NewGeneratePdfTask(payload)
	if err != nil {
		return err
	}
	s.workerClient.Enqueue(pdfTask)
	return nil
}
