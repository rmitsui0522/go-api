package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type User struct {
	ID          uint      `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	MailAddress string    `json:"mailAddress"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (h *handler) getUsers(c echo.Context) error {
	var users []User

	err := h.DB.Find(&users).Error

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}
