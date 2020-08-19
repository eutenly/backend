package authentication

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"golang.org/x/oauth2"
)

func SpotifyAuthenticationRoutes(e *echo.Echo) {
	e.GET("/login/spotify", func(c echo.Context) error {
		//Check if user logged in
		sess, _ := session.Get("session", c)
		if sess.Values["authed"] != true {
			return c.String(http.StatusUnauthorized, "You are not logged in. Please login to Eutenly before continuing.")
		}

		//Create login URL
		clientID := os.Getenv("SPOTIFY_CLIENT_ID")
		redirectURI := os.Getenv("WEBSERVER_URL") + "/auth/spotify"
		scopes := "user-read-playback-state user-modify-playback-state user-read-recently-played user-library-modify playlist-modify-private"
		authURL := fmt.Sprintf("https://accounts.spotify.com/authorize?response_type=code&client_id=%v&scope=%v&redirect_uri=%v", clientID, scopes, redirectURI)

		//Send user there
		return c.Redirect(http.StatusTemporaryRedirect, authURL)
	})
	e.GET("/auth/spotify", func(c echo.Context) error {
		//Get OAUTH2 login code
		authCode := c.QueryParam("code")

		//If no token was passed then error
		if authCode == "" {
			return c.String(http.StatusUnauthorized, "A login error occured.")
		}

		//Request accessToken
		accessToken, err := authenticateSpotify(authCode)
		if err != nil {
			return c.String(http.StatusUnauthorized, "A Spotify login error occured. "+err.Error())
		}

		return c.String(http.StatusOK, "access token: "+accessToken)
	})
}

func authenticateSpotify(authCode string) (accessToken string, returnErr error) {
	conf := &oauth2.Config{
		ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("WEBSERVER_URL") + "/auth/spotify",
		Scopes:       []string{"user-read-playback-state user-modify-playback-state user-read-recently-played user-library-modify playlist-modify-private"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.spotify.com/authorize",
			TokenURL: "https://accounts.spotify.com/api/token",
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
