package main

import (
	"go.uber.org/fx"
	"log"
	"simtix-ticketing/clients/amqp"
	"simtix-ticketing/config"
	"simtix-ticketing/database"
	"simtix-ticketing/worker"
	"simtix-ticketing/worker/handlers"
)

// consumer server
func startServer(server *worker.WorkerServer) {
	log.Print("Starting worker sever")
	server.Run()
}

func main() {
	app := fx.New(
		config.Module,
		database.Module,
		worker.Module,
		amqp.Module,
		handlers.Module,
		fx.Invoke(startServer),
		//fx.NopLogger,
	)
	app.Run()
}
