package routes

import (
	"echo-api/controllers"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo) {
	userController := controllers.NewUserController()
	usersGroup := e.Group("api/users")

	// Routes
	usersGroup.GET("/", userController.GetAll)
	usersGroup.GET("/:id", userController.GetById)
	usersGroup.POST("/create", userController.Create)
	usersGroup.PUT("/:id/update", userController.Update)
	usersGroup.DELETE("/:id/delete", userController.Delete)
}
