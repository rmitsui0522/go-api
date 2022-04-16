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

	e.GET(BASE_URL+"/health", h.health)

	e.GET(BASE_URL+"/users", h.getUsers)
	e.POST(BASE_URL+"/users", h.createUser)
	e.GET(BASE_URL+"/users/:id", h.getUser)
	e.PUT(BASE_URL+"/users/:id", h.updateUser)
	e.DELETE(BASE_URL+"/users/:id", h.deleteUser)


	return e
}

func (h *handler) health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "OK"})
}
