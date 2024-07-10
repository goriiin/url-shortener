package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func New(log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		log.Info("logger middleware initialized")

		fn := func(w http.ResponseWriter, r *http.Request) {

			entry := log.With(
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("remote_addr", r.RemoteAddr),
				slog.String("user_agent", r.UserAgent()),
				// TODO: когда добавлю
				//slog.String("request_id", r.Header.Get("X-Request-Id")),
			)

			next.ServeHTTP(w, r)
			//status := r.Response.StatusCode
			//defer func(Body io.ReadCloser) {
			//	_ = Body.Close()
			//}(r.Body)
			t1 := time.Now()
			defer func() {
				entry.Info("request_finished",
					slog.String("method", r.Method),
					// TODO : разобраться
					//slog.Int("status", status),
					slog.Duration("end", time.Since(t1)),
					//slog.String("bytes", )
				)
			}()

		}
		return http.HandlerFunc(fn)
	}
}
