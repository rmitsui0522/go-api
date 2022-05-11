package logger

import (
	"net/http"

	"go-api/pkg/logger"
	"go-api/pkg/utility"
)

func NewMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			defer func() {
				if err := recover(); err != nil {
					utility.RespondJSON(w, http.StatusInternalServerError, map[string]interface{}{
						"message": err,
					})
					logger.Panic("panic")
				}
			}()

			next.ServeHTTP(w, r)

			logger.RequestInfo("request", r)
		})
	}
}
