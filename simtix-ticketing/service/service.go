package service

import (
	"go.uber.org/fx"
	"simtix-ticketing/service/event"
)

var Module = fx.Module("service",
	fx.Options(
		fx.Provide(
			fx.Annotate(event.NewEventService, fx.As(new(event.EventService))),
		),
		//fx.Provide(),
	),
)