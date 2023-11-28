package handler

import (
	"go.uber.org/fx"
	"simtix-ticketing/handler/event"
	"simtix-ticketing/handler/seat"
)

var Module = fx.Module(
	"handler",
	fx.Options(
		fx.Provide(
			fx.Annotate(event.NewEventHandler, fx.As(new(event.EventHandler))),
		),
		fx.Provide(
			fx.Annotate(seat.NewSeatHandler, fx.As(new(seat.SeatHandler))),
		),
	),
)
