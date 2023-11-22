package invoice

import (
	"github.com/pkg/errors"
	"net/http"
	"simtix/utils"
)

var DbErrCreateInvoice = &utils.BaseError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected DB error when creating invoice"),
}

var ErrInvoiceNotFound = &utils.BaseError{
	StatusCode: http.StatusBadRequest,
	Err:        errors.New("Invoice not found"),
}

var DbErrUpdateInvoice = &utils.BaseError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected DB error when updating invoice"),
}

var DbErrGetInvoice = &utils.BaseError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Unexpected DB error when getting invoice"),
}

var ErrBookingExist = &utils.BaseError{
	StatusCode: http.StatusBadRequest,
	Err:        errors.New("Invoice for that booking has already existed"),
}
