package handler

import (
	"go-api/pkg/v1/auth"
	"go-api/pkg/v1/health"
	"go-api/pkg/v1/users"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	handler := mux.NewRouter()
	api := handler.PathPrefix("/api").Subrouter()
	v1 := api.PathPrefix("/v1").Subrouter()

	handler.HandleFunc("/health", health.Health()).Methods("GET")

	v1.HandleFunc("/auth", auth.Authentication()).Methods("POST")

	v1.HandleFunc("/users", users.GetAllUsers()).Methods("GET")
	v1.HandleFunc("/users", users.CreateUser()).Methods("POST")
	v1.HandleFunc("/users/{id}", users.GetUser()).Methods("GET")
	v1.HandleFunc("/users/{id}", users.UpdateUser()).Methods("PUT")
	v1.HandleFunc("/users/{id}", users.DeleteUser()).Methods("DELETE")

	return handler
}
