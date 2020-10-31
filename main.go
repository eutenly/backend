package main

import (
	"fmt"
	"os"

	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"

	"./database"
	"./influxdb"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

func main() {
	logrus.Info("Eutenly Web Services v0.1")

	//Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Warn("Cannot load .env file. Will use your bash's variables instead.")
	}

	//Load Sentry Error Tracker
	err = sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
	})
	if err != nil {
		logrus.Fatalf("Sentry cannot be started: %s", err)
	}

	//Connect to database
	err = database.Connect()
	if err != nil {
		logrus.Fatalf("Cannot connect to DB: " + err.Error())
	}

	//Connect to influxdb
	influxdb.Connect()

	//Create Echo Instance
	e := echo.New()
	e.HideBanner = true

	//Setup error logging for Echo
	e.Use(middleware.Recover())
	e.Use(sentryecho.New(sentryecho.Options{}))

	//Setup API Router
	apiRouter(e)

	//Setup static router
	staticRouter(e)

	//Enable Sessions
	sessionManager := session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET"))))
	e.Use(sessionManager)

	//Start server
	logrus.Info("Starting Echo...")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", os.Getenv("PORT"))))
}
