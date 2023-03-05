package errors

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserError struct {
	message string
}

func (e *UserError) Error() string {
	return e.message
}

func NewUserError(message string) *UserError {
	return &UserError{message: message}
}

func HandleError(c echo.Context, err error) error {
	if userErr, ok := err.(*UserError); ok {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": userErr.message})
	}

	return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
}
