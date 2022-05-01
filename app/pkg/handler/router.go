package handler

import (
	"go-api/pkg/v1/auth"
	"go-api/pkg/v1/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func v1(e *echo.Group) {
	r := e.Group("/restricted")

	e.POST("/auth", controller.Authentication())

	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: auth.HmacSecret(),
	}))

	r.GET("/users", controller.GetAllUsers())
	r.POST("/users", controller.CreateUser())
	r.GET("/users/:id", controller.GetUser())
	r.PUT("/users/:id", controller.UpdateUser())
	r.DELETE("/users/:id", controller.DeleteUser())

}
