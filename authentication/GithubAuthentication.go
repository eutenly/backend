package authentication

import (
	"context"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"golang.org/x/oauth2"
)

func GithubAuthenticationRoutes(e *echo.Echo) {
	oauthConfig := &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		Scopes:       []string{""},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}
	e.GET("/login/github", func(c echo.Context) error {
		//Check if user logged in
		sess, _ := session.Get("session", c)
		if sess.Values["authed"] != true {
			return c.String(http.StatusUnauthorized, "You are not logged in. Please login to Eutenly before continuing.")
		}

		//Send user to consent screen
		return c.Redirect(http.StatusTemporaryRedirect, oauthConfig.AuthCodeURL(""))
	})
	e.GET("/auth/github", func(c echo.Context) error {
		//Get OAUTH2 login code
		authCode := c.QueryParam("code")

		//If no token was passed then error
		if authCode == "" {
			return c.String(http.StatusUnauthorized, "A login error occured.")
		}

		//Request accessToken
		_, err := authenticateGitHub(authCode, oauthConfig)
		if err != nil {
			return c.String(http.StatusUnauthorized, "A Github login error occured. "+err.Error())
		}

		//Set auth cookie
		c.SetCookie(&http.Cookie{Name: "authed_with", Value: "github"})

		//Redirect
		return c.Redirect(302, "/connections")
	})
}

func authenticateGitHub(authCode string, conf *oauth2.Config) (accessToken string, returnErr error) {

	ctx := context.Background()
	tok, err := conf.Exchange(ctx, authCode)
	if err != nil {
		returnErr = err
		return
	}
	return tok.AccessToken, err

}
