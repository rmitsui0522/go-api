package handler

import (
	"net/http"

	"go-api/pkg/middleware/auth0"
	"go-api/pkg/middleware/cors"
	"go-api/pkg/middleware/logger"
	v1 "go-api/pkg/v1"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func New() http.Handler {
	router := mux.NewRouter()
	v1.New(router.PathPrefix("/api/v1").Subrouter())

	log := logger.NewLogger()

	// initializing middlewares
	corsMiddleware := cors.NewMiddleware()
	logMiddleware := logger.NewMiddleware(log)
	jwtMiddleware := auth0.NewMiddleware()

	return alice.New(corsMiddleware, logMiddleware, jwtMiddleware).Then(router)
}
