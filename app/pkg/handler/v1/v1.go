package v1

import (
	"go-api/pkg/middleware/auth0"

	"go-api/pkg/handler/v1/auth"
	"go-api/pkg/handler/v1/users"
	"go-api/pkg/middleware"

	"github.com/gorilla/mux"
)

func RegisterHandlerFunc(api *mux.Router) {
	v1 := api.PathPrefix("/api/v1").Subrouter()
	jwtMiddleware := middleware.JWT

	v1.Use(jwtMiddleware)

	v1.HandleFunc("/auth", auth.Authentication()).Methods("POST")

	v1.HandleFunc("/users", auth0.UseScope(users.GetAllUsers(), auth0.READABLE)).Methods("GET")
	v1.HandleFunc("/users", auth0.UseScope(users.CreateUser(), auth0.WRITEABLE)).Methods("POST")
	v1.HandleFunc("/users/{id}", auth0.UseScope(users.GetUser(), auth0.READABLE)).Methods("GET")
	v1.HandleFunc("/users/{id}", auth0.UseScope(users.UpdateUser(), auth0.WRITEABLE)).Methods("PUT")
	v1.HandleFunc("/users/{id}", auth0.UseScope(users.DeleteUser(), auth0.WRITEABLE)).Methods("DELETE")
}
