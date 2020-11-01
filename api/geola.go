package api

import (
	"../influxdb"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"os"
)

type geolaRequest struct {
	MemberCount string `json:"memberCount"`
}

func geolaRoutes(app *echo.Echo) {
	app.POST("/api/v1/geola/join", func(c echo.Context) error {
		if c.Request().Header.Get("Authorization") != os.Getenv("GEOLA_ENDPOINT_KEY") {
			return c.NoContent(401)
		}

		//Get body
		var body geolaRequest
		err := c.Bind(&body)
		if err != nil {
			logrus.Error("Cannot assign body", err)
			return c.String(500, err.Error())
		}

		influxdb.CollectStat("eutenland_join_leave", map[string]string{"type": "join"}, map[string]interface{}{"totalMembers": body.MemberCount})
		return c.NoContent(200)
	})
	app.POST("/api/v1/geola/leave", func(c echo.Context) error {
		if c.Request().Header.Get("Authorization") != os.Getenv("GEOLA_ENDPOINT_KEY") {
			return c.NoContent(401)
		}

		//Get body
		var body geolaRequest
		err := c.Bind(&body)
		if err != nil {
			logrus.Error("Cannot assign body", err)
			return c.String(500, err.Error())
		}

		influxdb.CollectStat("eutenland_join_leave", map[string]string{"type": "leave"}, map[string]interface{}{"totalMembers": body.MemberCount})
		return c.NoContent(200)
	})
}
