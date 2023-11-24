package main

import (
	"go.uber.org/fx"
	"simtix/domain"
	"simtix/handlers"
	"simtix/lib"
	"simtix/migrations"
	"simtix/routes"
	"simtix/server"
)

func startApp(server *server.Server, database *lib.Database) {
	migrations.Up(database.DB)
	server.Run()
}

func main() {
	app := fx.New(
		server.Module,
		lib.Module,
		domain.Module,
		handlers.Module,
		routes.Module,
		fx.Invoke(startApp),
	)
	app.Run()
}
