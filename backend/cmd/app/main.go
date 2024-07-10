package main

import (
	"fmt"
	config2 "github.com/goriiin/myapp/backend/db/postgres"
	"github.com/goriiin/myapp/backend/internal/app/middleware"
	"github.com/goriiin/myapp/backend/internal/config"
	"github.com/goriiin/myapp/backend/pkg/sl"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	cfg := config.MustLoad()

	// TODO: pretty logger?
	log := config.SetupLogger(cfg.Env)

	log.Info("starting server", slog.String("key", cfg.Env))
	log.Debug("Debug mode enabled")

	storage, err := config2.New()
	if err != nil {
		log.Error("error creating postgres repository", sl.Err(err))
		os.Exit(1)
	}
	//err = repository.RmTables()
	//if err != nil {
	//	log.Error("error removing tables", sl.Err(err))
	//}

	_ = storage

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Welcome to the homepage!")
	})

	mux.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		// Перенаправляем запрос на страницу профиля пользователя
		http.Redirect(w, r, "/profile", http.StatusFound)
	})

	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "This is your profile page!")
	})
	loggerMiddleware := middleware.New(log)
	middlewareRec := middleware.NewRec(log)
	_ = middlewareRec

	srv := &http.Server{
		Addr:         cfg.HTTPServer.Address + ":" + cfg.HTTPServer.Port,
		WriteTimeout: cfg.HTTPServer.Timeout,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.Timeout,
		Handler:      middlewareRec(loggerMiddleware(mux)),
	}

	path := "http://" + cfg.HTTPServer.Address + ":" + cfg.HTTPServer.Port
	log.Info("starting server", slog.String("key", cfg.Env), slog.String("path", path))

	err = srv.ListenAndServe()
	if err != nil {
		log.Error("error starting server", sl.Err(err))
		os.Exit(1)
	}
}
