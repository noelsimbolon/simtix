package event

import (
	"github.com/gin-gonic/gin"
	"simtix-ticketing/handler/event"
)

type EventRoute struct {
	eventHandler event.EventHandler
}

func NewEventRoute(handler event.EventHandler) *EventRoute {
	return &EventRoute{
		eventHandler: handler,
	}
}

func (s EventRoute) Setup(rg *gin.RouterGroup) {
	eventRoute := rg.Group("/events")

	eventRoute.GET("/", s.eventHandler.GetAllEvents)
	eventRoute.GET("/:eventID", s.eventHandler.GetEventByID)
}