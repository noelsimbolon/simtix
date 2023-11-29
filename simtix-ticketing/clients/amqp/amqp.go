package amqp

import "go.uber.org/fx"

var Module = fx.Module("worker", fx.Options(fx.Provide(NewAmqpClient)))
