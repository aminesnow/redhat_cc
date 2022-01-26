package setup

import (
	"context"
	"net/http"
	"time"

	recover "github.com/dre1080/recovr"
)

func setupGlobalMiddleware(handler http.Handler) http.Handler {
	recovery := recover.New()
	return recovery(timeoutMiddleware(handler, 5*time.Second))
}

func timeoutMiddleware(next http.Handler, duration time.Duration) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), duration)
		defer cancel()

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}
