package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func staticRouter(e *echo.Echo) {
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
        Root:   "static",
        Index: "index.html",
        Browse: true,
        HTML5:  true,
    }))
}