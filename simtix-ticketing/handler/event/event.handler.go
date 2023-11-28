package event

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simtix-ticketing/service/event"
)

type EventHandler interface {
	GetAllEvents(c *gin.Context)
	GetEventByID(c *gin.Context)
}

type EventHandlerImpl struct {
	service event.EventService
}

func NewEventHandler(service event.EventService) *EventHandlerImpl {
	return &EventHandlerImpl{
		service: service,
	}
}

func (e *EventHandlerImpl) GetAllEvents(c *gin.Context) {
	events, err := e.service.GetAllEvents()

	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"error": err.Err.Error()})
		return
	}

	// Example response body if the events table is empty
	/**
	{
	    "events": []
	}
	*/

	// Example response body if the events table is not empty
	/**
	{
	    "events": [
	        {
	            "ID": "d3407173-3984-460e-8429-327b878667ff",
	            "createdAt": "2023-11-28T02:11:26.911367+07:00",
	            "updatedAt": "2023-11-28T02:11:26.911367+07:00",
	            "deletedAt": null,
	            "eventName": "Event 1",
	            "seats": null
	        },
	        {
	            "ID": "d1ef8e44-dd8d-4002-9c5a-0520b3fefcfd",
	            "createdAt": "2023-11-28T02:11:26.911367+07:00",
	            "updatedAt": "2023-11-28T02:11:26.911367+07:00",
	            "deletedAt": null,
	            "eventName": "Event 2",
	            "seats": null
	        }
	    ]
	}
	*/

	c.JSON(http.StatusOK, events)
}

func (e *EventHandlerImpl) GetEventByID(c *gin.Context) {
	eventID := c.Param("eventID")
	ev, err := e.service.GetEventByID(eventID)

	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"error": err.Err.Error()})
		return
	}

	//

	c.JSON(http.StatusOK, ev)
}
