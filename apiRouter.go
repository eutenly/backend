package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func apiRouter(e *echo.Echo) {
	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Eutenly Backend")
	})
}
