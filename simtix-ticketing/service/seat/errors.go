package seat

import (
	"github.com/pkg/errors"
	"net/http"
	"simtix-ticketing/error"
)

var DbErrGetSeats = &error.Error{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected database error when getting seats."),
}

var DbErrGetSeat = &error.Error{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected database error when getting seat."),
}

var ErrSeatNotFound = &error.Error{
	StatusCode: http.StatusNotFound,
	Err:        errors.New("Seat not found."),
}

var ErrSeatExists = &error.Error{
	StatusCode: http.StatusBadRequest,
	Err:        errors.New("Seat already exist"),
}

var DbErrCreateSeat = &error.Error{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected database error when creating seat."),
}

var ErrSeatNotAvailable = &error.Error{
	StatusCode: http.StatusBadRequest,
	Err:        errors.New("Seat not available for booking."),
}

var ErrExternalCallFailed = &error.Error{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Failed to hold seat."),
}

var ErrMakePaymentFailed = &error.Error{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Failed to make payment."),
}

var DbErrUpdateSeat = &error.Error{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected database error when updating seat."),
}

var ErrEventNotExist = &error.Error{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Event doesn't exist"),
}
