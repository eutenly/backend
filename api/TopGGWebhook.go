package api

import (
	"encoding/json"
	"github.com/getsentry/sentry-go"
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
			sentry.CaptureException(err)
			return c.String(http.StatusBadRequest, "Bad request body")
		}

		//Parse body
		var data votehook
		err = json.Unmarshal(body, &data)

		//Store vote
		user, err := database.GetUser(data.User)
		if err != nil {
			sentry.CaptureException(err)
			return c.NoContent(http.StatusInternalServerError)
		}
		user.VoteExpireTimestamp = int32((time.Now().UnixNano() / 1000000) + ((12 * time.Hour).Milliseconds()))
		err = database.SetUser(user)
		if err != nil {
			sentry.CaptureException(err)
			return c.NoContent(http.StatusInternalServerError)
		}

		//Return
		return c.String(http.StatusOK, "OK")
	})
}
