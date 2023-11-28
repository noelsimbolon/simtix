package dto

type CreateEventDto struct {
	EventName string `json:"eventName" binding:"required"`
	EventTime string `json:"eventTime" binding:"required"`
}
