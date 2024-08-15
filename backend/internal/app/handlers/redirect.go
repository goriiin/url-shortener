package handlers

import (
	"github.com/goriiin/myapp/backend/internal/service"
	"log/slog"
	"net/http"
	"strings"
)

// TODO сделать структуру сервера
// TODO нужно отслеживать ошибки

func RedirectHandlerfunc(log *slog.Logger, serv service.URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.RedirectHandler"
		log = log.With(slog.String("op", op))
		path := r.URL.Path
		parts := strings.Split(path, "/")
		if len(parts) > 1 {
			alias := parts[1]
			u, err := serv.GetURL(alias)
			if err != nil {
				http.Error(w, "url no found", http.StatusNotFound)
			}
			http.Redirect(w, r, u.Url, http.StatusTemporaryRedirect)
		} else {
			http.Error(w, "alias Not Found", http.StatusNotFound)
		}

	}
}
