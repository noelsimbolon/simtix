package event

import (
	"github.com/pkg/errors"
	"net/http"
	"simtix-ticketing/error"
)

var DbErrGetAllEvents = &error.Error{
	StatusCode: http.StatusInternalServerError,
	Err: errors.New("Unexpected database error when getting all events."),
}

var DbErrGetEventById = &error.Error{
	StatusCode: http.StatusInternalServerError,
	Err: errors.New("Unexpected database error when getting an event by ID."),
}

var ErrEventNotFound = &error.Error{
	StatusCode: http.StatusBadRequest,
	Err: errors.New("Event not found with specified event ID."),
}
