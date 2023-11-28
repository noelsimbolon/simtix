package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	"os"
	"simtix/clients"
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
		domain.Module,
		handlers.Module,
		clients.Module,
		worker.Module,
		fx.Invoke(startServer),
		//fx.NopLogger,
	)
	app.Run()
	logger.Log.Info("oii")
}
