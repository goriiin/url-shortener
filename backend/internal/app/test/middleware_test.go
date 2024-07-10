package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestLoggingMiddleware(t *testing.T) {
	// Создаем тестовый обработчик, который возвращает ошибку 400.
	handler := loggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Bad Request")
	}))

	// Создаем тестовый запрос.
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	// Вызываем middleware.
	handler.ServeHTTP(w, req)

	// Проверяем код состояния ответа.
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}

	// Проверяем, что лог содержит информацию об ошибке 400.
	if !logsContain(t, "status=400") {
		t.Error("Log does not contain error status code 400")
	}
}

func logsContain(t *testing.T, message string) bool {
	t.Helper()
	return strings.Contains(fmt.Sprint(slog.Default().Handler()), message)
}
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		logger := slog.With(
			"method", r.Method,
			"path", r.URL.Path,
		)

		next.ServeHTTP(w, r)

		statusCode := w.Header().Get("Status")
		if statusCode == "" {
			statusCode = "200"
		}

		logger.Info("request completed",
			"duration", time.Since(start),
			"status", statusCode,
		)
	})
}
