package main

import (
	"echo-api/middleware"
	"echo-api/routes"

	"github.com/labstack/echo/v4"
)

const PORT = ":1323"

func main() {
	e := echo.New()

	// Middlewares
	middleware.SetMiddlewares(e)

	// Routes
	routes.NewRouter(e)

	e.Logger.Fatal(e.Start(PORT))
}
