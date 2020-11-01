package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func staticRouter(e *echo.Echo) {
	e.Use(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))

	e.Static("/", "static")
}
