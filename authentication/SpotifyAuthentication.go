package authentication

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

func SpotifyAuthenticationRoutes(e *echo.Echo) {
	oauthConfig := &oauth2.Config{
		ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("WEBSERVER_URL") + "/auth/spotify",
		Scopes:       []string{"user-read-private user-read-playback-state user-modify-playback-state user-read-recently-played user-library-modify playlist-modify-private playlist-modify-public user-top-read user-follow-modify"},
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

		//Get username
		username, id, err := loginSpotify(accessToken)
		if err != nil {
			logrus.Error("spotify login:", err)
			username = "?"
		}

		//Store tokens
		err = storeTokens(fmt.Sprint(sess.Values["discord_id"]), "spotify", id, username, map[string]string{"accessToken": accessToken, "refreshToken": refreshToken})
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

func loginSpotify(accessToken string) (username string, id string, err error) {
	//Construct body of request
	request, err := http.NewRequest("GET", "https://api.spotify.com/v1/me", bytes.NewBuffer([]byte("")))
	if err != nil {
		return
	}

	//Set custom UA & Auth because Reddit is annoying
	request.Header.Add("Authorization", "Bearer "+accessToken)
	request.Header.Set("User-Agent", "eutenly-backend/0.1")

	//Send request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	//Check status
	if response.StatusCode != 200 {
		err = fmt.Errorf("bad status code ", response.StatusCode)
		return
	}

	//Format JSON from response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	responseJSON := make(map[string]interface{})
	err = json.Unmarshal(body, &responseJSON)
	if err != nil {
		return
	}

	//Detect errors within response
	if responseJSON["display_name"] == nil {
		err = fmt.Errorf("bad response")
		return
	}
	if responseJSON["id"] == nil {
		err = fmt.Errorf("bad response")
		return
	}

	username = fmt.Sprint(responseJSON["display_name"])
	id = fmt.Sprint(responseJSON["id"])

	return
}
