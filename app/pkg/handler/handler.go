package handler

import (
	"go-api/pkg/middleware/auth0"
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

	v1.HandleFunc("/users", auth0.UseJWT(users.GetAllUsers())).Methods("GET")
	v1.HandleFunc("/users", auth0.UseJWT(users.CreateUser())).Methods("POST")
	v1.HandleFunc("/users/{id}", auth0.UseJWT(users.GetUser())).Methods("GET")
	v1.HandleFunc("/users/{id}", auth0.UseJWT(users.UpdateUser())).Methods("PUT")
	v1.HandleFunc("/users/{id}", auth0.UseJWT(users.DeleteUser())).Methods("DELETE")

	// auth0.GetAuth0Token()
	return handler
}
