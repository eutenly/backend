package authentication

import (
	"context"
	"fmt"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"net/http"
	"os"

	"github.com/maiacodes/fetch"
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
		RedirectURL: os.Getenv("WEBSERVER_URL") + "/auth/youtube",
	}
	e.GET("/login/youtube", func(c echo.Context) error {
		//Get session
		sess, _ := session.Get("session", c)
		if sess.Values["authed"] != true {
			return c.Redirect(302, "/login/discord?redirect_to=/login/youtube")
		}

		redirect := fmt.Sprintf("https://accounts.google.com/o/oauth2/v2/auth?client_id=%v&redirect_uri=%v/auth/youtube&scope=https://www.googleapis.com/auth/youtube.readonly&response_type=code&access_type=offline", os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("WEBSERVER_URL"))

		return c.Redirect(http.StatusFound, redirect)
	})
	e.GET("/auth/youtube", func(c echo.Context) error {
		//Get OAUTH2 login code
		authCode := c.QueryParam("code")

		//If no token was passed then error
		if authCode == "" {
			return loginError(fmt.Errorf("no auth code"), "youtube", c)
		}

		//Get session
		sess, _ := session.Get("session", c)
		if sess.Values["authed"] != true {
			return c.Redirect(302, "/login/discord?redirect_to=/login/google")
		}

		// Exchange token
		tok, err := oauthConfig.Exchange(context.TODO(), c.QueryParam("code"))
		if err != nil {
			return loginError(err, "youtube", c)
		}

		// Get user info
		var r ytResp
		err = fetch.FetchJSON("https://www.googleapis.com/youtube/v3/channels?part=snippet&mine=true", "GET", nil, &r, fetch.FetchOptions{Authorization: "Bearer " + tok.AccessToken})
		if err != nil {
			return loginError(err, "youtube", c)
		}
		username := "Unknown"
		id := "unknown"
		if len(r.Items) != 0 {
			id = r.Items[0].ID
			username = r.Items[0].Snippet.Title
		}

		//Store tokens
		err = storeTokens(fmt.Sprint(sess.Values["discord_id"]), "youtube", id, username, map[string]string{"accessToken": tok.AccessToken, "refreshToken": tok.RefreshToken})
		if err != nil {
			return loginError(err, "youtube", c)
		}

		//Set auth cookie
		authCookie("youtube", c)

		//Redirect
		return c.Redirect(302, "/connections")
	})
	return
}

type ytResp struct {
	Items []ytChan `json:"items"`
}

type ytChan struct {
	ID      string `json:"id"`
	Snippet ytSnip `json:"snippet"`
}

type ytSnip struct {
	Title string `json:"title"`
}
