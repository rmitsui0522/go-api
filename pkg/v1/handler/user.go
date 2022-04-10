package handler

import (
	"net/http"
	"time"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type User struct {
	ID          uint      `json:"id" param:"id"`
	FirstName   string    `json:"firstName" validate:"required"`
	LastName    string    `json:"lastName" validate:"required"`
	MailAddress string    `json:"mailAddress" validate:"required,email"`
	CreateAt    time.Time `json:"createAt"`
	UpdateAt    time.Time `json:"updateAt"`
}

func (h *handler) getUsers(c echo.Context) error {
	var users []User

	err := h.DB.Find(&users).Error

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func (h *handler) createUser(c echo.Context) error {
	var user User
	validate := validator.New()

	user.CreateAt = time.Now().Round(time.Second)
	user.UpdateAt = time.Now().Round(time.Second)

	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(&user); err != nil {
		return c.String(http.StatusNotAcceptable, err.Error())
	}

	if err := h.DB.Create(&user).Error; err != nil {
		return c.String(http.StatusNotAcceptable, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *handler) getUser(c echo.Context) error {
	var user User
	paramId := c.Param("id")

	err := h.DB.Where("id=?", paramId).Find(&user).Error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (h *handler) updateUser(c echo.Context) error {
	var user User
	var data User
	validate := validator.New()

	paramId := c.Param("id")
	data.UpdateAt = time.Now().Round(time.Second)

	if err := c.Bind(&data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(&user); err != nil {
		return c.String(http.StatusNotAcceptable, err.Error())
	}

	err := h.DB.Where("id=?", paramId).Find(&user).Update(&data).Error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (h *handler) deleteUser(c echo.Context) error {
	var user User
	paramId := c.Param("id")

	err := h.DB.Where("id=?", paramId).Delete(&user).Error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
