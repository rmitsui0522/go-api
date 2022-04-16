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
	auth := api.Group("/auth")
	users := api.Group("/users")

	auth.POST("/", h.auth)

	users.GET("/", h.getUsers)
	users.POST("/", h.createUser)
	users.GET("/:id", h.getUser)
	users.PUT("/:id", h.updateUser)
	users.DELETE("/:id", h.deleteUser)

	return e
}

func (h *handler) health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "OK"})
}
