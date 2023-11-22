package domain

import (
	"go.uber.org/fx"
	"simtix/domain/invoice"
)

var Module = fx.Module(
	"domain",
	fx.Options(
		fx.Provide(
			fx.Annotate(invoice.NewInvoiceService, fx.As(new(invoice.InvoiceService))),
		),
	),
)
