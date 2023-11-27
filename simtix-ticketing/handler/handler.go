package handler

import (
	"go.uber.org/fx"
	"simtix-ticketing/handler/event"
)

var Module = fx.Module("handler",
	fx.Options(
		fx.Provide(
			fx.Annotate(event.NewEventHandler, fx.As(new(event.EventHandler))),
		),
	),
)