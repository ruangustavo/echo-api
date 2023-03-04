package main

import (
	"echo-api/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	middleware.SetMiddlewares(e)
}
