package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

func Health() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "OK"})
	}
}
