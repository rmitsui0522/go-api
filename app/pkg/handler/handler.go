package handler

import (
	"net/http"

	"go-api/pkg/v1/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const BASE_URL = "/api"

func New() http.Handler {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/health", controller.Health())

	v1(e.Group(BASE_URL + "/v1"))

	return e
}
