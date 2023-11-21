package main

import (
	"go.uber.org/fx"
	"main/domain"
	"main/handlers"
	"main/lib"
	"main/migrations"
	"main/routes"
	"main/server"
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
