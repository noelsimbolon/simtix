package route

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"simtix-ticketing/route/event"
	"simtix-ticketing/route/seat"
)

var Module = fx.Module(
	"route",
	fx.Options(
		fx.Provide(event.NewEventRoute),
		fx.Provide(seat.NewSeatRoute),
		fx.Provide(NewRoute),
	),
)

type Routes []Route

type Route interface {
	Setup(rg *gin.RouterGroup)
}

func NewRoute(eventRoute *event.EventRoute, seatRoute *seat.SeatRoute) *Routes {
	return &Routes{
		eventRoute,
		seatRoute,
	}
}

func (r Routes) Setup(rg *gin.RouterGroup) {
	for _, route := range r {
		route.Setup(rg)
	}
}
