package handler

import (
	"net/http"

	"go-api/pkg/handler/health"
	v1 "go-api/pkg/handler/v1"
	"go-api/pkg/middleware"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	// initialize routers
	api := mux.NewRouter()

	// initialize middlewares
	corsMiddleware := middleware.CORS
	logMiddleware := middleware.Logger

	// apply middlewares
	api.Use(corsMiddleware)
	api.Use(logMiddleware)

	// register routes
	api.HandleFunc("/health", health.Health()).Methods("GET")
	v1.RegisterHandlerFunc(api)

	return api
}
