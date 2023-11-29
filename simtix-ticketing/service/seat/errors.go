package seat

import (
	"github.com/pkg/errors"
	"net/http"
	"simtix-ticketing/utils"
)

var DbErrGetSeats = &utils.Error{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected database error when getting seats."),
}

var DbErrGetSeat = &utils.Error{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected database error when getting seat."),
}

var ErrSeatNotFound = &utils.Error{
	StatusCode: http.StatusNotFound,
	Err:        errors.New("Seat not found."),
}

var ErrSeatExists = &utils.Error{
	StatusCode: http.StatusBadRequest,
	Err:        errors.New("Seat already exist"),
}

var DbErrCreateSeat = &utils.Error{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected database error when creating seat."),
}

var ErrSeatNotAvailable = &utils.Error{
	StatusCode: http.StatusBadRequest,
	Err:        errors.New("Seat not available for booking."),
}

var ErrExternalCallFailed = &utils.Error{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Failed to hold seat."),
}

var ErrMakePaymentFailed = &utils.Error{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Failed to make payment."),
}

var DbErrUpdateSeat = &utils.Error{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected database error when updating seat."),
}

var ErrEventNotExist = &utils.Error{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Event doesn't exist"),
}
