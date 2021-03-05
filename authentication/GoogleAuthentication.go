package authentication

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"golang.org/x/oauth2"
	"net/http"
	"os"
)

func GoogleAuthenticationRoutes(e *echo.Echo) {
	oauthConfig := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/youtube.readonly"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/v2/auth",
			TokenURL: "https://oauth2.googleapis.com/token",
		},
		RedirectURL: os.Getenv("WEBSERVER_URL") + "/auth/google",
	}
	e.GET("/login/google", func(c echo.Context) error {
		//Get session
		sess, _ := session.Get("session", c)
		if sess.Values["authed"] != true {
			return c.Redirect(302, "/login/discord?redirect_to=/login/google")
		}

		redirect := fmt.Sprintf("https://accounts.google.com/o/oauth2/v2/auth?client_id=%v&redirect_uri=%v/auth/google&scope=https://www.googleapis.com/auth/youtube.readonly&response_type=code", os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("WEBSERVER_URL"))

		return c.Redirect(http.StatusFound, redirect)
	})
	e.GET("/auth/google", func(c echo.Context) error {
		//Get OAUTH2 login code
		authCode := c.QueryParam("code")

		//If no token was passed then error
		if authCode == "" {
			return loginError(fmt.Errorf("no auth code"), "github", c)
		}

		//Get session
		sess, _ := session.Get("session", c)
		if sess.Values["authed"] != true {
			return c.Redirect(302, "/login/discord?redirect_to=/login/google")
		}

		// Exchange token
		tok, err := oauthConfig.Exchange(context.TODO(), c.QueryParam("code"))
		if err != nil {
			return loginError(err, "google", c)
		}

		//Store tokens
		err = storeTokens(fmt.Sprint(sess.Values["discord_id"]), "google", "", "", map[string]string{"accessToken": tok.AccessToken})
		if err != nil {
			return loginError(err, "github", c)
		}

		//Set auth cookie
		authCookie("google", c)

		//Redirect
		return c.Redirect(302, "/connections")
	})
	return
}
