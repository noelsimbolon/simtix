package main

import (
	"go.uber.org/fx"
	"simtix/lib"
	"simtix/server"
	"simtix/worker"
)

// consumer server
func startServer(server *worker.WorkerServer) {
	server.Run()
}

func main() {
	app := fx.New(
		server.Module,
		lib.Module,
		worker.Module,
		fx.Invoke(startServer),
	)
	app.Run()
}
