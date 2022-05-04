package logger

import (
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

func NewMiddleware(logger zerolog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					logger.Error().Interface("err", err).Msg("recovered")
				}
			}()

			next.ServeHTTP(w, r)

			start := time.Now()
			logger.Debug().Str("method", r.Method).Str("path", r.URL.EscapedPath()).Interface("duration", time.Since(start)).Msg("requested")
		})
	}
}
