package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dghubble/oauth1"
	"github.com/dghubble/oauth1/twitter"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

func twitterAuthenticationRoutes(e *echo.Echo) {
	config := oauth1.Config{
		ConsumerKey:    os.Getenv("TWITTER_KEY"),
		ConsumerSecret: os.Getenv("TWITTER_SECRET"),
		CallbackURL:    os.Getenv("WEBSERVER_URL") + "/auth/twitter",
		Endpoint:       twitter.AuthorizeEndpoint,
	}

	e.GET("/login/twitter", func(c echo.Context) error {
		sess, _ := session.Get("session", c)

		//Check if user is authed with Discord
		if sess.Values["authed"] == false {
			return c.String(http.StatusInternalServerError, "You need to login.")
		}

		//Make Twitter Auth Request
		requestToken, requestSecret, err := config.RequestToken()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		//Store values
		sess.Values["twitter_requestToken"] = requestToken
		sess.Values["twitter_requestSecret"] = requestSecret
		sess.Save(c.Request(), c.Response())

		//Generate Auth URL and Redirect
		authorizationURL, err := config.AuthorizationURL(requestToken)
		return c.Redirect(http.StatusTemporaryRedirect, authorizationURL.String())
	})

	e.GET("/auth/twitter", func(c echo.Context) error {
		sess, _ := session.Get("session", c)

		//Check if user is authed with Discord
		if sess.Values["authed"] == false {
			return c.String(http.StatusInternalServerError, "You need to login.")
		}

		//Get Auth Values
		oauthVerifier := c.QueryParam("oauth_verifier")
		requestToken := fmt.Sprint(sess.Values["twitter_requestToken"])
		requestSecret := fmt.Sprint(sess.Values["twitter_requestSecret"])

		//Get Twitter Access Tokens
		accessToken, accessSecret, err := config.AccessToken(requestToken, requestSecret, oauthVerifier)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusAccepted, accessSecret+" "+accessToken)
	})

}
