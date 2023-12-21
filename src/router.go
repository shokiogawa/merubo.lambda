package main

import "github.com/labstack/echo/v4"

func NewRouter(init *Initialize) (e *echo.Echo) {
	e = echo.New()
	e.POST("/api/contact", init.LineHandler.SendContactMessageEcho)
	return
}
