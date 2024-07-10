package middleware

import (
	uuid "github.com/satori/go.uuid"
	"log/slog"
	"net/http"
)

// TODO: проверка на аунтификацию

// TODO: RequestID
func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.NewV4()
		r.Header.Set("X-Request-ID", id.String())
		w.Header().Set("X-Request-ID", id.String())
		next.ServeHTTP(w, r)
	})
}

// NewRec TODO: Recover - восстановка паники - нужен тест
func NewRec(log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					log.Error("panic recovered", "err", err)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

// TODO URLFormat
