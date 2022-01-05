package controllers

import (
	"net/http"
	"technical-test-aino-golang/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func FetchAllUser(c echo.Context) error {
	result, err := models.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func Register(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	// run StoreUser function on user model
	result, err := models.StoreUser(name, email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func Login(c echo.Context) error {
	// validate request input
	v := validator.New()
	email := c.FormValue("email")
	err := v.Var(email, "required,email")
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	password := c.FormValue("password")
	err = v.Var(password, "required")
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// run login function on user model
	result, err := models.Login(email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Status:  http.StatusInternalServerError,
			Message: "Email atau password salah!",
		})
	}

	return c.JSON(http.StatusOK, result)
}
