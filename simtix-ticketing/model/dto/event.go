package dto

type CreateEventDto struct {
	EventName string `json:"eventName" binding:"required"`
}
