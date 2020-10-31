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

func SpotifyAuthenticationRoutes(e *echo.Echo) {
	oauthConfig := &oauth2.Config{
		ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("WEBSERVER_URL") + "/auth/spotify",
		Scopes:       []string{"user-read-playback-state user-modify-playback-state user-read-recently-played user-library-modify playlist-modify-private playlist-modify-public user-top-read"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.spotify.com/authorize",
			TokenURL: "https://accounts.spotify.com/api/token",
		}}
	e.GET("/login/spotify", func(c echo.Context) error {
		//Check if user logged in
		sess, _ := session.Get("session", c)
		if sess.Values["authed"] != true {
			return c.Redirect(302, "/login/discord?redirect_to=/login/spotify")
		}

		//Send user to consent screen
		return c.Redirect(http.StatusTemporaryRedirect, oauthConfig.AuthCodeURL(""))
	})
	e.GET("/auth/spotify", func(c echo.Context) error {
		//Get OAUTH2 login code
		authCode := c.QueryParam("code")

		//If no token was passed then error
		if authCode == "" {
			return c.String(http.StatusUnauthorized, "No key was passed.")
		}

		//Request accessToken
		accessToken, refreshToken, err := authenticateSpotify(authCode, oauthConfig)
		if err != nil {
			return loginError(err, "spotify", c)
		}

		//Get session
		sess, _ := session.Get("session", c)
		if sess.Values["authed"] != true {
			return c.Redirect(302, "/login/discord?redirect_to=/login/spotify")
		}

		//Store tokens
		err = storeTokens(fmt.Sprint(sess.Values["discord_id"]), "spotify", "123", map[string]string{"accessToken": accessToken, "refreshToken": refreshToken})
		if err != nil {
			return loginError(err, "spotify", c)
		}

		//Set auth cookie
		c.SetCookie(&http.Cookie{Name: "authed_with", Value: "spotify"})

		//Redirect
		return c.Redirect(302, "/connections")
	})
}

func authenticateSpotify(authCode string, conf *oauth2.Config) (accessToken string, refreshToken string, requestError error) {

	ctx := context.Background()
	tok, err := conf.Exchange(ctx, authCode)
	if err != nil {
		requestError = err
		return
	}
	return tok.AccessToken, tok.RefreshToken, err

}
