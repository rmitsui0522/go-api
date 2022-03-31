package handler

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(db *gorm.DB) http.Handler {
	e := echo.New()
	e.Use(middleware.CORS())

	h := &handler{
		DB: db,
	}

	e.GET("/health", h.health)

	e.GET("/users", h.getUsers)
	e.POST("/users", h.createUser)
	e.GET("/users/:id", h.getUser)
	e.PUT("/users/:id", h.updateUser)
	e.DELETE("/users/:id", h.deleteUser)

	return e
}

type handler struct {
	DB *gorm.DB
}

func (h *handler) health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "OK"})
}
