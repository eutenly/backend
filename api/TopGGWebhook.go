package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"../database"

	"github.com/labstack/echo"
)

type votehook struct {
	User string `json:"user"`
}

func TopGGWebhook(app *echo.Echo) {

	app.POST("/api/topgg/votehook", func(c echo.Context) error {

		//Validate request
		webhookSecret := c.Request().Header.Get("Authorization")
		if webhookSecret != os.Getenv("TOPGG_WEBHOOK_SECRET") {
			return c.String(http.StatusUnauthorized, "Invalid webhook secret")
		}

		//Get body
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			return c.String(http.StatusBadRequest, "Bad request body")
		}

		//Parse body
		var data votehook
		err = json.Unmarshal(body, &data)

		//Set vote expire timestamp
		voteExpireTimestamp := (time.Now().UnixNano() / 1000000) + ((12 * time.Hour).Milliseconds())
		database.FindOneAndUpdate("users", map[string]interface{}{"_id": data.User}, map[string]interface{}{"voteExpireTimestamp": voteExpireTimestamp}, true)

		//Return
		return c.String(http.StatusOK, "OK")
	})
}
