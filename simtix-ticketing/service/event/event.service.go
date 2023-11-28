package event

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"simtix-ticketing/config"
	"simtix-ticketing/database"
	"simtix-ticketing/error"
	"simtix-ticketing/model/event"
	"time"
)

type EventService interface {
	GetAllEvents() (*event.GetAllEventsDao, *error.Error)
	GetEventByID(eventID string) (*event.Event, *error.Error)
	CreateEvent(dto *event.CreateEventDto) (*event.Event, *error.Error)
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

func (s *EventServiceImpl) GetAllEvents() (*event.GetAllEventsDao, *error.Error) {
	var events []event.Event

	err := s.repository.Find(&events).Error

	if err != nil {
		return nil, DbErrGetAllEvents
	}

	return &event.GetAllEventsDao{
		Events: events,
	}, nil
}

func (s *EventServiceImpl) GetEventByID(eventID string) (*event.Event, *error.Error) {
	var ev event.Event

	err := s.repository.Where("id = ?", eventID).Preload("Seats", "event_id = ?", eventID).First(&ev).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrEventNotFound
		}

		return nil, DbErrGetEventById
	}

	return &ev, nil
}

func (s *EventServiceImpl) CreateEvent(dto *event.CreateEventDto) (*event.Event, *error.Error) {
	eventTime, err := time.Parse(time.RFC822, dto.EventTime)
	if err != nil {
		log.Print(err)
		return nil, ErrInvalidTime
	}
	event := event.Event{
		EventName: dto.EventName,
		EventTime: eventTime,
	}
	err = s.repository.Create(&event).Error
	if err != nil {
		return nil, DbErrCreateEvent
	}
	return &event, nil
}
