package middleware

import (
	"net/http"

	"go-api/pkg/middleware/auth0"
	"go-api/pkg/middleware/cors"
	"go-api/pkg/middleware/logger"
)

var CORS func(http.Handler) http.Handler
var Logger func(http.Handler) http.Handler
var JWT func(http.Handler) http.Handler

func init() {
	CORS = cors.NewMiddleware()
	Logger = logger.NewMiddleware()
	JWT = auth0.NewMiddleware()
}
