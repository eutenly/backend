package main

import (
	"./api"
	"./authentication"

	"github.com/labstack/echo"
)

func apiRouter(app *echo.Echo) {
	//myEutenly Routes
	api.MyEutenlyRoutes(app)

	//Authentication Routes
	authentication.DiscordAuthenticationRoutes(app)
	authentication.GithubAuthenticationRoutes(app)
	authentication.TwitterAuthenticationRoutes(app)
	authentication.SpotifyAuthenticationRoutes(app)

	//top.gg webhook
	api.TopGGWebhook(app)
}
