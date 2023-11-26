package clients

import (
	"go.uber.org/fx"
	"simtix/clients/ticketing"
)

var Module = fx.Module("clients", fx.Provide(ticketing.NewTicketingClient))
