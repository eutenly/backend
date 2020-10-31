package api

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"

	"../database"
	"../database/schemas"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

func deleteConnection(e *echo.Echo) {
	e.DELETE("/api/v1/connection", func(c echo.Context) error {
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

		//Get body
		var body map[string]interface{} = make(map[string]interface{})
		err = c.Bind(&body)
		if err != nil {
			logrus.Error(err)
			return c.String(500, err.Error())
		}

		//If connection exists, remove it
		if _, ok := user.Connections[fmt.Sprint(body["connection"])]; ok {
			delete(user.Connections, fmt.Sprint(body["connection"]))
		} else {
			return c.NoContent(500)
		}

		//Due to weird bug with Mongo Struct decoder, if there's
		//no connections left, just unset the object.
		if len(user.Connections) == 0 {
			err = database.UnsetUserConnection(*user.ID)
			if err != nil {
				logrus.Error(err)
				return c.String(500, err.Error())
			}
			return c.NoContent(200)
		}

		//Save user
		err = database.SetUser(user)
		if err != nil {
			logrus.Error(err)
			return c.String(500, err.Error())
		}

		//Respond
		return c.NoContent(200)
	})
}
