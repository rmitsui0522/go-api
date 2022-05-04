package handler

import (
	"net/http"

	"go-api/pkg/middleware/auth0"
	"go-api/pkg/middleware/logger"
	"go-api/pkg/v1/controller"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	handler := mux.NewRouter()
	api := handler.PathPrefix("/api").Subrouter()
	v1 := api.PathPrefix("/v1").Subrouter()

	loggerMiddleware := logger.NewMiddleware()
	jwtMiddleware := auth0.NewMiddleware()

	handler.Use(loggerMiddleware)
	v1.Use(jwtMiddleware)

	handler.HandleFunc("/health", controller.Health()).Methods("GET")

	v1.HandleFunc("/auth", controller.Authentication()).Methods("POST")

	v1.HandleFunc("/users", controller.GetAllUsers()).Methods("GET")
	v1.HandleFunc("/users", controller.CreateUser()).Methods("POST")
	v1.HandleFunc("/users/{id}", controller.GetUser()).Methods("GET")
	v1.HandleFunc("/users/{id}", controller.UpdateUser()).Methods("PUT")
	v1.HandleFunc("/users/{id}", controller.DeleteUser()).Methods("DELETE")

	// r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey: auth.HmacSecret(),
	// }))

	// e.GET("/users", echo.WrapHandler(auth0.UseJWT(controller.GetAllUsers())))

	return handler
}
