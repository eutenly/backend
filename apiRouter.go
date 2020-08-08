package main

import (
	"github.com/labstack/echo"
)

func apiRouter(app *echo.Echo) {
	//myEutenly Routes
	myEutenlyRoutes(app)

	//Authentication Routes
	discordAuthenticationRoutes(app)
	githubAuthenticationRoutes(app)
	twitterAuthenticationRoutes(app)

	//top.gg webhook
	topggWebhook(app)
}
