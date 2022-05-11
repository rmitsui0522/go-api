package handler

import (
	"net/http"

	"go-api/pkg/handler/health"
	v1 "go-api/pkg/handler/v1"
	"go-api/pkg/middleware"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func New() http.Handler {
	router := mux.NewRouter()
	v1.New(router.PathPrefix("/api/v1").Subrouter())

	router.HandleFunc("/health", health.Health()).Methods("GET")

	// initializing middlewares
	corsMiddleware := middleware.CORS
	logMiddleware := middleware.Logger
	jwtMiddleware := middleware.JWT

	return alice.New(corsMiddleware, logMiddleware, jwtMiddleware).Then(router)
}
