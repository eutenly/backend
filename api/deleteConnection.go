package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"

	"eutenly/backend/database"
	"eutenly/backend/database/schemas"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func deleteConnection(e *echo.Echo) {
	e.DELETE("/api/v1/connection", func(c echo.Context) error {
		//Get session
		sess, _ := session.Get("session", c)
		if sess.Values["authed"] != true {
			return c.String(http.StatusUnauthorized, "Not logged in")
		}

		//Fetch account info
		userid := fmt.Sprint(sess.Values["discord_id"])
		user, err := database.GetUser(userid)
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
		connection := fmt.Sprint(body["connection"])

		//If connection exists, remove it
		if _, ok := user.Connections[connection]; ok {
			delete(user.Connections, connection)
		} else {
			return c.NoContent(500)
		}

		//Send manual `uncache` request to bot
		cache := uncacheConnection(connection, userid)
		if cache != nil {
			logrus.Error(cache)
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

func uncacheConnection(connection string, userid string) (err error) {
	//Construct body of request
	query, err := json.Marshal(map[string]interface{}{
		"user_id":    userid,
		"connection": connection,
	})
	request, err := http.NewRequest("POST", fmt.Sprintf("%v/api/v1/uncacheConnection", os.Getenv("BOT_ENDPOINT")), bytes.NewBuffer(query))
	if err != nil {
		return err
	}

	//Headers
	request.Header.Set("User-Agent", "eutenly-backend/0.1")
	request.Header.Set("Authorization", os.Getenv("BOT_ENDPOINT_KEY"))

	//Send request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	//Format JSON from response body
	if response.StatusCode != 200 {
		return fmt.Errorf("bad status: %v", response.StatusCode)
	}

	return
}
