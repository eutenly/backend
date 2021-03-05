package main

import (
	"eutenly/backend/api"
	"eutenly/backend/authentication"
	"github.com/labstack/echo/v4/middleware"
	"net/url"
	"os"

	"github.com/labstack/echo/v4"
)

func apiRouter(app *echo.Echo) {
	//myEutenly Routes
	api.MyEutenlyRoutes(app)

	//Authentication Routes
	authentication.DiscordAuthenticationRoutes(app)
	authentication.GithubAuthenticationRoutes(app)
	authentication.TwitterAuthenticationRoutes(app)
	authentication.SpotifyAuthenticationRoutes(app)
	authentication.RedditAuthenticationRoutes(app)
	authentication.GoogleAuthenticationRoutes(app)

	//top.gg webhook
	api.TopGGWebhook(app)

	if os.Getenv("development") == "true" {
		url1, _ := url.Parse("http://localhost:8000")
		targets := []*middleware.ProxyTarget{
			{
				URL: url1,
			},
		}
		app.Group("*").Use(middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))

	}
}
