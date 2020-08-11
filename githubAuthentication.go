package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"golang.org/x/oauth2"
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
		//Get OAUTH2 login code
		authCode := c.QueryParam("code")

		//If no token was passed then error
		if authCode == "" {
			return c.String(http.StatusUnauthorized, "A Discord login error occured.")
		}

		//Request accessToken
		accessToken, err := authenticateGitHub(authCode)
		if err != nil {
			return c.String(http.StatusUnauthorized, "A Github login error occured. "+err.Error())
		}

		return c.String(http.StatusOK, "access token: "+accessToken)
	})
}

func authenticateGitHub(authCode string) (accessToken string, returnErr error) {
	conf := &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		Scopes:       []string{""},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}

	ctx := context.Background()
	tok, err := conf.Exchange(ctx, authCode)
	if err != nil {
		returnErr = err
		return
	}
	return tok.AccessToken, err

}
