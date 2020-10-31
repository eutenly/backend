package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func MyEutenlyRoutes(e *echo.Echo) {
	e.GET("/api/v1", func(c echo.Context) error {
		return c.String(http.StatusOK, "MyEutenly API")
	})
	myConnections(e)
	deleteConnection(e)
	geolaRoutes(e)
}
