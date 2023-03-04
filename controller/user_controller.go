package controller

import (
	"echo-api/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) GetAll(c echo.Context) error {
	users := models.GetAllUsers()
	return c.JSON(http.StatusOK, users)
}

func (u *UserController) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "It wasn't possible convert string id to integer"})
	}

	user, err := models.GetUserById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (u *UserController) Create(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		c.String(http.StatusBadRequest, "Not working")
	}

	if err := models.CreateUser(user); err != nil {
		c.String(http.StatusInternalServerError, "Not working")
	}

	return c.JSON(http.StatusCreated, user)
}
