package controllers

import (
	"echo-api/controllers/errors"
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
	user, err := u.getUserFromContext(c)

	if err != nil {
		return errors.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (u *UserController) Create(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		errors.HandleError(c, errors.NewUserError("Invalid user data"))
	}

	if err := models.CreateUser(user); err != nil {
		return errors.HandleError(c, errors.NewUserError("Failed to create user"))
	}

	return c.JSON(http.StatusCreated, user)
}

func (u *UserController) Update(c echo.Context) error {
	user, err := u.getUserFromContext(c)

	if err != nil {
		return errors.HandleError(c, errors.NewUserError("Invalid user data"))
	}

	if err := models.UpdateUser(user); err != nil {
		return errors.HandleError(c, errors.NewUserError("Failed to update user"))
	}

	return c.JSON(http.StatusOK, user)
}

func (u *UserController) Delete(c echo.Context) error {
	user, err := u.getUserFromContext(c)

	if err != nil {
		return errors.HandleError(c, err)
	}

	if err := models.DeleteUser(user); err != nil {
		return errors.HandleError(c, errors.NewUserError("Failed to delete user"))
	}

	return c.JSON(http.StatusOK, user)
}

func (u *UserController) getUserFromContext(c echo.Context) (*models.User, error) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return nil, errors.NewUserError("Invalid user ID")
	}

	user, err := models.GetUserById(id)

	if err != nil {
		return nil, errors.NewUserError("User not found")
	}

	return user, nil
}
