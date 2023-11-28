package service

import (
	"go.uber.org/fx"
	"simtix-ticketing/service/event"
	"simtix-ticketing/service/seat"
)

var Module = fx.Module(
	"service",
	fx.Options(
		fx.Provide(
			fx.Annotate(event.NewEventService, fx.As(new(event.EventService))),
		),
		fx.Provide(
			fx.Annotate(seat.NewSeatService, fx.As(new(seat.SeatService))),
		),
		//fx.Provide(),
	),
)
