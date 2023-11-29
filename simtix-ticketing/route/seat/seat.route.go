package seat

import (
	"github.com/gin-gonic/gin"
	"simtix-ticketing/handler/seat"
)

type SeatRoute struct {
	seatHandler seat.SeatHandler
}

func NewSeatRoute(handler seat.SeatHandler) *SeatRoute {
	return &SeatRoute{
		seatHandler: handler,
	}
}

func (r SeatRoute) Setup(rg *gin.RouterGroup) {
	seatRoute := rg.Group("/seat")
	seatRoute.GET("/", r.seatHandler.GetSeatsByEvent)
	seatRoute.GET("/:seatID", r.seatHandler.GetSeatByID)
	seatRoute.POST("/", r.seatHandler.PostSeat)
	seatRoute.PATCH("/", r.seatHandler.PatchSeatForBooking)
	seatRoute.POST("/webhook", r.seatHandler.SeatWebhook)
}
