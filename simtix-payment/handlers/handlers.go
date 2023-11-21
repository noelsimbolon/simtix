package handlers

import "go.uber.org/fx"

var Module = fx.Module(
	"invoice",
	fx.Options(
		fx.Provide(
			fx.Annotate(NewInvoiceHandlerImpl, fx.As(new(InvoiceHandler))),
		),
	),
)
