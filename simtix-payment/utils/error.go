package utils

import "net/http"

type BaseError struct {
	StatusCode int
	Err        error
}

type InternalServerError struct {
	BaseError
}

func NewInternalServerError(err error) *InternalServerError {
	return &InternalServerError{
		BaseError: BaseError{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		},
	}
}

type NotFoundError struct {
	BaseError
}

func NewNotFoundError(err error) *NotFoundError {
	return &NotFoundError{
		BaseError: BaseError{
			StatusCode: http.StatusNotFound,
			Err:        err,
		},
	}
}
