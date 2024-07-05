package main

import (
	"github.com/goriiin/myapp/backend/internal/config"
	"github.com/goriiin/myapp/backend/internal/libs/sl"
	"github.com/goriiin/myapp/backend/internal/storage/postgres"
	"log/slog"
	"os"
)

func main() {
	cfg := config.MustLoad()

	log := config.SetupLogger(cfg.Env)

	log.Info("starting server", slog.String("key", cfg.Env))
	log.Debug("Debug mode enabled")

	storage, err := postgres.New()
	if err != nil {
		log.Error("error creating postgres storage", sl.Err(err))
		os.Exit(1)
	}
	//err = storage.RmTables()
	//if err != nil {
	//	log.Error("error removing tables", sl.Err(err))
	//}

	_ = storage

	// TODO: init router: standard, render
	path := "http://" + cfg.HTTPServer.Address + ":" + cfg.HTTPServer.Port
	log.Info("starting server", slog.String("key", cfg.Env), slog.String("path", path))
	// TODO: run server
}
