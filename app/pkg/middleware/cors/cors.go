package cors

import (
	"net/http"
	"strings"
)

func NewMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			if AccessControlAllowOrigin(origin) {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}
			if r.Method == "OPTIONS" {
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				w.Header().Set("Access-Control-Allow-Methods", "GET,POST")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization")
				return
			} else {
				next.ServeHTTP(w, r)
			}
		})
	}
}

func AccessControlAllowOrigin(requestOrigin string) bool {
	allowedOrigins := []string{"localhost:3000"}
	for _, allowedOrigin := range allowedOrigins {
		if strings.Contains(requestOrigin, allowedOrigin) {
			return true
		}
	}
	return false
}
