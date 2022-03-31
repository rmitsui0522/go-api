package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type User struct {
	ID          uint      `json:"id" param:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	MailAddress string    `json:"mailAddress"`
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
	err := c.Bind(&user)
	if err != nil {
		return err
	}

	err = h.DB.Create(&user).Error
	if err != nil {
		return err
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
	paramId := c.Param("id")
	data.UpdateAt = time.Now().Round(time.Second)

	err := c.Bind(&data)
	if err != nil {
		return c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
	}

	err = h.DB.Where("id=?", paramId).Find(&user).Update(&data).Error
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
