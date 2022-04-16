package handler

import (
	"net/http"
	"strconv"

	"go-api/pkg/v1/model"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

func (h *handler) getUsers(c echo.Context) error {
	users, err := model.FindUsers()

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, users)
}

func (h *handler) createUser(c echo.Context) error {
	var user model.User
	validate := validator.New()

	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(&user); err != nil {
		return c.String(http.StatusNotAcceptable, err.Error())
	}

	if err := model.CreateUser(&user); err != nil {
		return c.String(http.StatusNotAcceptable, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *handler) getUser(c echo.Context) error {
	paramId := c.Param("id")
	id, _ := strconv.ParseUint(paramId, 10, 64)
	user, err := model.FindUser(&model.User{ID: uint(id)})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (h *handler) updateUser(c echo.Context) error {
	var data model.User
	validate := validator.New()

	paramId := c.Param("id")
	id, _ := strconv.ParseUint(paramId, 10, 64)

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	if err := validate.Struct(&data); err != nil {
		return c.JSON(http.StatusNotAcceptable, map[string]string{"message": err.Error()})
	}

	user, err := model.UpdateUser(&model.User{ID: uint(id)}, &data)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (h *handler) deleteUser(c echo.Context) error {
	paramId := c.Param("id")
	id, _ := strconv.ParseUint(paramId, 10, 64)

	user, err := model.DeleteUser(&model.User{ID: uint(id)})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
