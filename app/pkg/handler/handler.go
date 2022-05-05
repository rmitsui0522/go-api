package handler

import (
	"go-api/pkg/middleware/auth0"
	"go-api/pkg/v1/controller"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	handler := mux.NewRouter()
	api := handler.PathPrefix("/api").Subrouter()
	v1 := api.PathPrefix("/v1").Subrouter()

	handler.HandleFunc("/health", controller.Health()).Methods("GET")

	v1.HandleFunc("/auth", controller.Authentication()).Methods("POST")

	v1.HandleFunc("/users", controller.GetAllUsers()).Methods("GET")
	v1.HandleFunc("/users", controller.CreateUser()).Methods("POST")
	v1.HandleFunc("/users/{id}", controller.GetUser()).Methods("GET")
	v1.HandleFunc("/users/{id}", controller.UpdateUser()).Methods("PUT")
	v1.HandleFunc("/users/{id}", controller.DeleteUser()).Methods("DELETE")

	return handler
}
