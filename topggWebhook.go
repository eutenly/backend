package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func topggWebhook(app *echo.Echo) {
	app.POST("/api/topgg/votehook", func(c echo.Context) error {
		// if c.Request().Header
		return c.String(http.StatusServiceUnavailable, "Service Unavailable")
	})
}
