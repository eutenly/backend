package main

import (
	"github.com/labstack/echo"
)

func apiRouter(app *echo.Echo) {
	//myEutenly Routes
	myEutenlyRoutes(app)

	//Discord API Routes
	discordAuthenticationRoutes(app)

	//top.gg webhook
	topggWebhook(app)
}
