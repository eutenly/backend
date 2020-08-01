package main

import (
	"github.com/labstack/echo"
)

func staticRouter(e *echo.Echo) {
	e.Static("/", "static");
}