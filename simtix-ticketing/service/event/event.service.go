package event

import (
	"errors"
	"gorm.io/gorm"
	"simtix-ticketing/config"
	"simtix-ticketing/database"
	"simtix-ticketing/error"
	"simtix-ticketing/model/event"
)

type EventService interface {
	GetAllEvents() (*event.GetAllEventsDao, *error.Error)
	GetEventByID(eventID string) (*event.Event, *error.Error)
}

type EventServiceImpl struct {
	config *config.Config
	repository *gorm.DB
}

func NewEventService(config *config.Config, database *database.Database) *EventServiceImpl {
	return &EventServiceImpl{
		config: config,
		repository: database.DB,
	}
}

func (s *EventServiceImpl) GetAllEvents() (*event.GetAllEventsDao, *error.Error) {
	var events []event.Event

	result := s.repository.Find(&events)

	if result.Error != nil {
		return nil, DbErrGetAllEvents
	}

	return &event.GetAllEventsDao{
		Events: events,
	}, nil
}

func (s *EventServiceImpl) GetEventByID(eventID string) (*event.Event, *error.Error) {
	var ev event.Event

	result := s.repository.Where("id = ?", eventID).Preload("Seats", "event_id = ?", eventID).First(&ev)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrEventNotFound
		}

		return nil, DbErrGetEventById
	}

	return &ev, nil
}
