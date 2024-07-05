package main

import (
	"github.com/goriiin/myapp/backend/internal/config"
	"log/slog"
)

func main() {
	cfg := config.MustLoad()

	log := config.SetupLogger(cfg.Env)

	log.Info("starting server", slog.String("key", cfg.Env))
	log.Debug("Debug mode enabled")

	// TODO: init logger: slog

	// TODO: init storage: Postgres

	// TODO: init router: standard, render
	path := "http://" + cfg.HTTPServer.Address + ":" + cfg.HTTPServer.Port
	log.Info("starting server", slog.String("key", cfg.Env), slog.String("path", path))
	// TODO: run server
}
