package event

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"simtix-ticketing/config"
	"simtix-ticketing/database"
	"simtix-ticketing/model"
	"simtix-ticketing/model/dao"
	"simtix-ticketing/model/dto"
	"simtix-ticketing/utils"
	"time"
)

type EventService interface {
	GetAllEvents() (*dao.GetAllEventsDao, *utils.Error)
	GetEventByID(eventID string) (*model.Event, *utils.Error)
	CreateEvent(dto *dto.CreateEventDto) (*model.Event, *utils.Error)
}

type EventServiceImpl struct {
	config     *config.Config
	repository *gorm.DB
}

func NewEventService(config *config.Config, database *database.Database) *EventServiceImpl {
	return &EventServiceImpl{
		config:     config,
		repository: database.DB,
	}
}

func (s *EventServiceImpl) GetAllEvents() (*dao.GetAllEventsDao, *utils.Error) {
	var events []model.Event

	err := s.repository.Find(&events).Error

	if err != nil {
		return nil, DbErrGetAllEvents
	}

	return &dao.GetAllEventsDao{
		Events: events,
	}, nil
}

func (s *EventServiceImpl) GetEventByID(eventID string) (*model.Event, *utils.Error) {
	var ev model.Event

	err := s.repository.Where("id = ?", eventID).Preload("Seats", "event_id = ?", eventID).First(&ev).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrEventNotFound
		}

		return nil, DbErrGetEventById
	}

	return &ev, nil
}

func (s *EventServiceImpl) CreateEvent(dto *dto.CreateEventDto) (*model.Event, *utils.Error) {
	eventTime, err := time.Parse(time.RFC822, dto.EventTime)
	if err != nil {
		log.Print(err)
		return nil, ErrInvalidTime
	}
	event := model.Event{
		EventName: dto.EventName,
		EventTime: eventTime,
	}
	err = s.repository.Create(&event).Error
	if err != nil {
		return nil, DbErrCreateEvent
	}
	return &event, nil
}
