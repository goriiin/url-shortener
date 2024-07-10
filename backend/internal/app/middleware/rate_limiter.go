package middleware

import (
	"net/http"
	"sync"
	"time"
)

// TODO: Rate limiting middleware
type rateLimiter struct {
	limit  int
	window time.Duration
	mu     sync.Mutex
	tokens chan struct{}
}

// rl := newRateLimiter(10, time.Minute)
// http.Handle("/ping", rl.Middleware(http.DefaultServeMux))
func newRateLimiter(limit int, window time.Duration) *rateLimiter {
	return &rateLimiter{
		limit:  limit,
		window: window,
		tokens: make(chan struct{}, limit),
	}
}

func (rl *rateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rl.mu.Lock()
		select {
		case rl.tokens <- struct{}{}:
			// Allow request
		default:
			// Rate limit exceeded
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		rl.mu.Unlock()

		go func() {
			time.Sleep(rl.window)
			rl.mu.Lock()
			<-rl.tokens
			rl.mu.Unlock()
		}()

		next.ServeHTTP(w, r)
	})
}
