package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

func githubAuthenticationRoutes(e *echo.Echo) {
	e.GET("/login/github", func(c echo.Context) error {
		//Check if user logged in
		sess, _ := session.Get("session", c)
		if sess.Values["authed"] != true {
			return c.String(http.StatusUnauthorized, "You are not logged in. Please login to Eutenly before continuing.")
		}

		//Create login URL
		clientID := os.Getenv("GITHUB_CLIENT_ID")
		redirectURI := os.Getenv("WEBSERVER_URL") + "/auth/github"
		authURL := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%v&redirect_uri=%v", clientID, redirectURI)

		//Send user there
		return c.Redirect(http.StatusTemporaryRedirect, authURL)
	})
	e.GET("/auth/github", func(c echo.Context) error {
		return c.String(http.StatusForbidden, "not ready")
	})
}
