package v1

import (
	"go-api/pkg/middleware/auth0"

	"go-api/pkg/v1/auth"
	"go-api/pkg/v1/health"
	"go-api/pkg/v1/users"

	"github.com/gorilla/mux"
)

func New(v1 *mux.Router) {
	v1.HandleFunc("/health", health.Health()).Methods("GET")

	v1.HandleFunc("/auth", auth.Authentication()).Methods("POST")

	v1.HandleFunc("/users", auth0.UseScope(users.GetAllUsers(), auth0.READABLE)).Methods("GET")
	v1.HandleFunc("/users", auth0.UseScope(users.CreateUser(), auth0.WRITEABLE)).Methods("POST")
	v1.HandleFunc("/users/{id}", auth0.UseScope(users.GetUser(), auth0.READABLE)).Methods("GET")
	v1.HandleFunc("/users/{id}", auth0.UseScope(users.UpdateUser(), auth0.WRITEABLE)).Methods("PUT")
	v1.HandleFunc("/users/{id}", auth0.UseScope(users.DeleteUser(), auth0.WRITEABLE)).Methods("DELETE")
}
