package httpClient

import "go.uber.org/fx"

var Module = fx.Module("httpClient", fx.Provide(NewPaymentClient))
