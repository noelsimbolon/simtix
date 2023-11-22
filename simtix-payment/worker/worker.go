package worker

import "go.uber.org/fx"

var Module = fx.Module("worker", fx.Options(fx.Provide(NewWorkerClient), fx.Provide(NewServer)))
