package event

import (
	"github.com/pkg/errors"
	"net/http"
	"simtix-ticketing/utils"
)

var DbErrGetAllEvents = &utils.Error{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected database error when getting all events."),
}

var DbErrGetEventById = &utils.Error{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected database error when getting an event by ID."),
}

var ErrEventNotFound = &utils.Error{
	StatusCode: http.StatusBadRequest,
	Err:        errors.New("Event not found with specified event ID."),
}

var ErrInvalidTime = &utils.Error{
	StatusCode: http.StatusBadRequest,
	Err:        errors.New("Invalid event time given. Time should be in format RFC822: 02 Sep 15 08:00 WIB"),
}

var DbErrCreateEvent = &utils.Error{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected database error when creating event."),
}
