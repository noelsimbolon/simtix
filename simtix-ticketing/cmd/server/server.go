package main

import (
	"go.uber.org/fx"
	"simtix-ticketing/clients/httpClient"
	"simtix-ticketing/config"
	"simtix-ticketing/database"
	"simtix-ticketing/handler"
	"simtix-ticketing/migration"
	"simtix-ticketing/route"
	"simtix-ticketing/server"
	"simtix-ticketing/service"
	"simtix-ticketing/worker"
)

func startApp(server *server.Server, database *database.Database) {
	migration.Up(database.DB)
	server.Run()
}

func main() {
	app := fx.New(
		server.Module,
		config.Module,
		database.Module,
		service.Module,
		handler.Module,
		worker.Module,
		route.Module,
		httpClient.Module,
		fx.Invoke(startApp),
		fx.NopLogger,
	)

	app.Run()
}
