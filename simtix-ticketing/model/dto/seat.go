package dto

import "simtix-ticketing/model"

type CreateSeatDto struct {
	EventID string `json:"eventID" binding:"required"`
	Status model.SeatStatus `json:"status" binding:"required"`
}
