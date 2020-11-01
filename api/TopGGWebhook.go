package api

import (
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
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

	app.POST("/api/v1/vote", func(c echo.Context) error {

		//Validate request
		webhookSecret := c.Request().Header.Get("Authorization")
		if webhookSecret != os.Getenv("TOPGG_WEBHOOK_SECRET") {
			return c.String(http.StatusUnauthorized, "Invalid webhook secret")
		}

		//Get body
		var body votehook
		err := c.Bind(&body)
		if err != nil {
			logrus.Error(err)
			return c.String(http.StatusBadRequest, err.Error())
		}

		//Store vote
		user, err := database.GetUser(body.User)
		if err != nil {
			sentry.CaptureException(err)
			return c.NoContent(http.StatusInternalServerError)
		}
		user.VoteExpireTimestamp = int(time.Now().Add(time.Hour*12).UnixNano() / int64(time.Millisecond))
		err = database.SetUser(user)
		if err != nil {
			sentry.CaptureException(err)
			return c.NoContent(http.StatusInternalServerError)
		}

		//Return
		return c.String(http.StatusOK, "OK")
	})
}
