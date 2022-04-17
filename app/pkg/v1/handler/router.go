package handler

import (
	"go-api/pkg/v1/controller"

	"github.com/labstack/echo"
)

func initRouting(e *echo.Echo) {
	e.GET("/health", controller.Health())

	api := e.Group(BASE_URL)

	api.POST("/auth", controller.Auth())

	api.GET("/users", controller.GetAllUsers())
	api.POST("/users", controller.CreateUser())
	api.GET("/users/:id", controller.GetUser())
	api.PUT("/users/:id", controller.UpdateUser())
	api.DELETE("/users/:id", controller.DeleteUser())
}
