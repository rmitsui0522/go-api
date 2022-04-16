package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const BASE_URL = "/api/v1"

type handler struct{}

func New() http.Handler {
	e := echo.New()
	e.Use(middleware.CORS())

	h := &handler{}

	e.GET("/health", h.health)

	api := e.Group(BASE_URL)

	api.POST("/auth", h.auth)

	api.GET("/users", h.getUsers)
	api.POST("/users", h.createUser)
	api.GET("/users/:id", h.getUser)
	api.PUT("/users/:id", h.updateUser)
	api.DELETE("/users/:id", h.deleteUser)

	return e
}

func (h *handler) health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "OK"})
}
