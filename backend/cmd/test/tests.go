package main

import (
	"context"
	"net/http"
	"strings"
)

type params struct {
	ID string
}

func paramMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		parts := strings.Split(path, "/")
		if len(parts) < 4 {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}
		id := parts[2]
		ctx := r.Context()
		ctx = context.WithValue(ctx, "id", id)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/create/{id}/today", func(w http.ResponseWriter, r *http.Request) {
		id, ok := r.Context().Value("id").(string)
		if !ok {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		w.Write([]byte("Hello, " + id))
	})

	http.ListenAndServe(":8080", paramMiddleware(mux))
}
