package main

import (
	"github.com/goriiin/myapp/backend/internal/app/handlers"
	"github.com/goriiin/myapp/backend/internal/app/middleware"
	"github.com/goriiin/myapp/backend/internal/config"
	"github.com/goriiin/myapp/backend/internal/repository/postgres"
	service "github.com/goriiin/myapp/backend/internal/service"
	"github.com/goriiin/myapp/backend/pkg/sl"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	cfg := config.MustLoad()

	log := config.SetupLogger(cfg.Env)

	log.Info("starting server", slog.String("key", cfg.Env))
	log.Debug("Debug mode enabled")

	storage := postgres.New()

	serv := service.NewUrlSaverService(storage)

	mux := http.NewServeMux()
	mux.HandleFunc("/{alias}", func(w http.ResponseWriter, r *http.Request) {
		handlers.RedirectHandlerfunc(log, serv)(w, r)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.Get(log, serv)(w, r)
		case http.MethodPost:
			handlers.Save(log, serv)(w, r)
		case http.MethodPut:
			handlers.Edit(log, serv)(w, r)
		case http.MethodDelete:
			handlers.Delete(log, serv)(w, r)
		}
	})

	loggerMiddleware := middleware.New(log)
	middlewareRec := middleware.RecoverMiddleware(log)

	srv := &http.Server{
		Addr:         cfg.HTTPServer.Address + ":" + cfg.HTTPServer.Port,
		WriteTimeout: cfg.HTTPServer.Timeout,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.Timeout,
		Handler:      middlewareRec(loggerMiddleware(mux)),
	}

	path := "http://" + cfg.HTTPServer.Address + ":" + cfg.HTTPServer.Port
	log.Info("starting server", slog.String("key", cfg.Env), slog.String("path", path))

	err := srv.ListenAndServe()
	if err != nil {
		log.Error("error starting server", sl.Err(err))
		os.Exit(1)
	}
}
