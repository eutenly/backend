package api

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"

	"eutenly/backend/database"
	"eutenly/backend/database/schemas"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type connectionResp struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	ConnectedAt int    `json:"connectedAt"`
}

type accountDetailsResp struct {
	Username    string                    `json:"username"`
	Discrim     string                    `json:"discrim"`
	ID          string                    `json:"id"`
	Connections map[string]connectionResp `json:"connections"`
}

func myConnections(e *echo.Echo) {
	e.GET("/api/v1/me", func(c echo.Context) error {
		if os.Getenv("development") == "true" {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		}
		//Get session
		sess, _ := session.Get("session", c)
		if sess.Values["authed"] != true {
			return c.String(http.StatusUnauthorized, "Not logged in")
		}

		//Fetch account info
		user, err := database.GetUser(fmt.Sprint(sess.Values["discord_id"]))
		if err != nil {
			logrus.Error(err)
			return c.String(500, err.Error())
		}
		if user.Connections == nil {
			user.Connections = make(map[string]schemas.Connection)
		}

		//Construct response
		response := accountDetailsResp{}
		response.ID = fmt.Sprint(sess.Values["discord_id"])
		response.Username = fmt.Sprint(sess.Values["discord_username"])
		response.Discrim = fmt.Sprint(sess.Values["discord_discrim"])

		//Connections
		response.Connections = make(map[string]connectionResp)
		for service, connection := range user.Connections {
			response.Connections[service] = connectionResp{
				ID:          connection.ID,
				Username:    connection.Username,
				ConnectedAt: connection.ConnectedAt,
			}
		}

		//Serve
		return c.JSON(http.StatusOK, response)
	})
}
