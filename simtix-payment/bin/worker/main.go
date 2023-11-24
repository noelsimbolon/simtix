package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	"os"
	"simtix/domain"
	"simtix/lib"
	"simtix/utils/logger"
	"simtix/worker"
	"simtix/worker/handlers"
)

// consumer server
func startServer(server *worker.WorkerServer) {
	logger.Log.Info("Starting worker server")
	server.Run()
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	app := fx.New(
		lib.Module,
		worker.Module,
		domain.Module,
		handlers.Module,
		fx.Invoke(startServer),
		//fx.NopLogger,
	)
	app.Run()
	logger.Log.Info("oii")
}
