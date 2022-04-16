package handler

import (
	"net/http"

	"go-api/pkg/v1/model"

	"github.com/labstack/echo"
)

type AuthRequest struct {
	Account  string `json:"mailAddress" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	success bool
	user    model.User
}

func (h *handler) auth(c echo.Context) error {
	var req AuthRequest
	var res AuthResponse
	if err := c.Bind(&req); err != nil {
		return err
	}

	user, err := model.FindUser(&model.User{Account: req.Account})
	if err != nil {
		return err
	}

	if user.ID == 0 || user.Password != req.Password {
		return c.String(http.StatusUnauthorized, "Invalid MailAddress or Password.")
	}

	// session := session.Default(c)
	// session.Set("loginCompleted", "completed")
	// session.Save()

	res.success = true
	res.user = user

	return c.JSON(http.StatusOK, user)
}
