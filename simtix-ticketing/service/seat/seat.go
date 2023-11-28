package seat

import (
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"log"
	"simtix-ticketing/clients/httpClient"
	"simtix-ticketing/config"
	"simtix-ticketing/database"
	"simtix-ticketing/error"
	event2 "simtix-ticketing/model/event"
	"simtix-ticketing/model/seat"
)

type SeatService interface {
	GetSeatsByEventID(eventID string) ([]seat.Seat, *error.Error)
	GetSeatByID(id string) (*seat.Seat, *error.Error)
	CreateSeat(dto seat.CreateSeatDto) (*seat.Seat, *error.Error)
	BookSeat(updateSeatStatusDto seat.BookSeatDto) (
		*seat.Seat, *error.Error,
	)
}

type SeatServiceImpl struct {
	config        *config.Config
	repository    *gorm.DB
	paymentClient *httpClient.PaymentClient
}

func NewSeatService(
	config *config.Config, database *database.Database, client *httpClient.PaymentClient,
) *SeatServiceImpl {
	return &SeatServiceImpl{
		config:        config,
		repository:    database.DB,
		paymentClient: client,
	}
}

func (s *SeatServiceImpl) GetSeatsByEventID(eventID string) ([]seat.Seat, *error.Error) {
	var seats []seat.Seat
	err := s.repository.Where("event_id = ?", eventID).Find(&seats).Error
	if err != nil {
		return nil, DbErrGetSeats
	}
	return seats, nil
}

func (s *SeatServiceImpl) GetSeatByID(id string) (*seat.Seat, *error.Error) {
	var seat seat.Seat
	err := s.repository.Where("id = ?", id).First(&seat).Error
	if err != nil {
		return nil, DbErrGetSeat
	}
	return &seat, nil
}

func (s *SeatServiceImpl) CreateSeat(seatDto seat.CreateSeatDto) (*seat.Seat, *error.Error) {
	var event event2.Event
	dbErr := s.repository.Where("id = ?", seatDto.EventID).First(&event).Error
	if dbErr != nil {
		return nil, ErrEventNotExist
	}

	var existingSeat seat.Seat
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

	seat := seat.Seat{
		EventID:    seatDto.EventID,
		SeatNumber: seatDto.SeatNumber,
		SeatRow:    seatDto.SeatRow,
		Status:     seat.SEATSTATUS_OPEN,
		Price:      seatDto.Price,
	}
	err = s.repository.Create(&seat).Error
	if err != nil {
		return nil, DbErrCreateSeat
	}

	return &seat, nil
}

// for webhook
func (s *SeatServiceImpl) UpdateSeatStatus(updateSeatStatusDto seat.UpdateSeatStatusDto) (*seat.Seat, *error.Error) {
	existingSeat, err := s.GetSeatByID(updateSeatStatusDto.SeatID)
	if err != nil {
		return nil, err
	}

	existingSeat.Status = updateSeatStatusDto.Status

	dbErr := s.repository.Save(&existingSeat).Error
	if dbErr != nil {
		return nil, DbErrCreateSeat
	}

	return existingSeat, nil
}

// for client service
func (s *SeatServiceImpl) BookSeat(updateSeatStatusDto seat.BookSeatDto) (
	*seat.Seat, *error.Error,
) {
	existingSeat, err := s.GetSeatByID(updateSeatStatusDto.SeatID)
	if err != nil {
		return nil, err
	}
	if existingSeat.Status != seat.SEATSTATUS_OPEN {
		return nil, ErrSeatNotAvailable
	}
	if !s.checkSeatExternally(updateSeatStatusDto.SeatID) {
		return nil, ErrExternalCallFailed
	}
	_, err = s.makeInvoice(existingSeat.Price, updateSeatStatusDto.BookingID)
	if err != nil {
		return nil, err
	}
	// what to do with invoice?
	existingSeat.Status = seat.SEATSTATUS_ONGOING
	dbErr := s.repository.Save(&existingSeat).Error
	if dbErr != nil {
		return nil, DbErrCreateSeat
	}
	return existingSeat, nil
}

var checkSeatExternallyCnt = 0

func (s *SeatServiceImpl) checkSeatExternally(seatID string) bool {
	checkSeatExternallyCnt += 1
	return checkSeatExternallyCnt%5 != 0
}

func (s *SeatServiceImpl) makeInvoice(amount decimal.Decimal, bookingID string) (
	*httpClient.PostInvoiceResponse, *error.Error,
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
