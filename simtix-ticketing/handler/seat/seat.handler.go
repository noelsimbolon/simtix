package seat

import (
	"github.com/gin-gonic/gin"
	"net/http"
	seat2 "simtix-ticketing/model/seat"
	"simtix-ticketing/service/seat"
)

type SeatHandler interface {
	GetSeatsByEvent(c *gin.Context)
	GetSeatByID(c *gin.Context)
	PostSeat(c *gin.Context)
	//PatchSeatStatus(c *gin.Context)
	PatchSeatForBooking(c *gin.Context)
}

type SeatHandlerImpl struct {
	service seat.SeatService
}

func NewSeatHandler(service seat.SeatService) *SeatHandlerImpl {
	return &SeatHandlerImpl{
		service: service,
	}
}

func (h *SeatHandlerImpl) GetSeatsByEvent(c *gin.Context) {
	eventID := c.Query("eventID")
	if eventID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Event ID parameter is required"})
		return
	}
	seats, err := h.service.GetSeatsByEventID(eventID)
	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"error": err.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, seats)
}

func (h *SeatHandlerImpl) GetSeatByID(c *gin.Context) {
	seatID := c.Param("seatID")
	if seatID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Seat ID parameter is required"})
		return
	}
	seat, err := h.service.GetSeatByID(seatID)
	if err != nil {
		c.AbortWithStatusJSON(err.StatusCode, gin.H{"error": err.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, seat)
	return
}

func (h *SeatHandlerImpl) PostSeat(c *gin.Context) {
	var dto seat2.CreateSeatDto
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seat, custErr := h.service.CreateSeat(dto)
	if custErr != nil {
		c.AbortWithStatusJSON(custErr.StatusCode, gin.H{"error": custErr.Err.Error()})
		return
	}

	c.JSON(http.StatusCreated, seat)
	return
}

func (h *SeatHandlerImpl) PatchSeatForBooking(c *gin.Context) {
	var dto seat2.BookSeatDto
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seat, custErr := h.service.BookSeat(dto)
	if custErr != nil {
		c.AbortWithStatusJSON(custErr.StatusCode, gin.H{"error": custErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, seat)
	return
}
