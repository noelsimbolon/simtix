package domain

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"domain",
	fx.Options(
		fx.Provide(
			fx.Annotate(NewInvoiceService, fx.As(new(InvoiceService))),
		),
	),
)
