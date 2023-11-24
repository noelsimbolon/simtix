package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	"os"
	"simtix/domain"
	"simtix/handlers"
	"simtix/lib"
	"simtix/migrations"
	"simtix/routes"
	"simtix/server"
	"simtix/worker"
)

func startApp(server *server.Server, database *lib.Database) {
	migrations.Up(database.DB)
	server.Run()
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	app := fx.New(
		server.Module,
		lib.Module,
		domain.Module,
		handlers.Module,
		routes.Module,
		worker.Module,
		fx.Invoke(startApp),
		//fx.NopLogger,
	)
	app.Run()
}
