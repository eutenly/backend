package api

import (
	"../influxdb"
	"github.com/labstack/echo"
	"os"
)

func geolaRoutes(app *echo.Echo) {
	app.POST("/api/v1/geola/join", func(c echo.Context) error {
		if c.Request().Header.Get("Authorization") != os.Getenv("GEOLA_ENDPOINT_KEY") {
			return c.NoContent(401)
		}
		influxdb.CollectStat("eutenland_join_leave", map[string]string{"type": "join"})
		return c.NoContent(200)
	})
	app.POST("/api/v1/geola/leave", func(c echo.Context) error {
		if c.Request().Header.Get("Authorization") != os.Getenv("GEOLA_ENDPOINT_KEY") {
			return c.NoContent(401)
		}
		influxdb.CollectStat("eutenland_join_leave", map[string]string{"type": "leave"})
		return c.NoContent(200)
	})
}
