package authentication

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo-contrib/session"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"golang.org/x/oauth2"
)

func RedditAuthenticationRoutes(e *echo.Echo) {
	oauthConfig := &oauth2.Config{
		ClientID:     os.Getenv("REDDIT_CLIENT_ID"),
		ClientSecret: os.Getenv("REDDIT_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("WEBSERVER_URL") + "/auth/reddit",
	}

	e.GET("/login/reddit", func(c echo.Context) error {
		//Check if user logged in
		sess, _ := session.Get("session", c)
		if sess.Values["authed"] != true {
			return c.Redirect(302, "/login/discord?redirect_to=/login/reddit")
		}

		//Construct custom url due to weird params
		authUrl := fmt.Sprintf("https://reddit.com/api/v1/authorize?client_id=%v&response_type=code&redirect_uri=%v&state=eutenly&duration=permanent&scope=identity,read,vote,save,history", oauthConfig.ClientID, oauthConfig.RedirectURL)

		//Send user to consent screen
		return c.Redirect(http.StatusTemporaryRedirect, authUrl)
	})
	e.GET("/auth/reddit", func(c echo.Context) error {
		//Get OAUTH2 login code
		authCode := c.QueryParam("code")

		//If no token was passed then error
		if authCode == "" {
			return c.String(http.StatusUnauthorized, "A login error occurred.")
		}

		//Request accessToken
		accessToken, refreshToken, err := authenticateReddit(authCode, oauthConfig)
		if err != nil {
			return loginError(err, "reddit", c)
		}

		//Get session
		sess, _ := session.Get("session", c)
		if sess.Values["authed"] != true {
			return c.Redirect(302, "/login/discord?redirect_to=/login/reddit")
		}

		//Store tokens
		err = storeTokens(fmt.Sprint(sess.Values["discord_id"]), "reddit", "123", "", map[string]string{"accessToken": accessToken, "refreshToken": refreshToken})
		if err != nil {
			return loginError(err, "reddit", c)
		}

		//Set auth cookie
		c.SetCookie(&http.Cookie{Name: "authed_with", Value: "reddit"})

		//Redirect
		return c.Redirect(302, "/connections")
	})
}

func authenticateReddit(authCode string, conf *oauth2.Config) (accessToken string, refreshToken string, err error) {
	//Construct body of request
	query := []byte(`grant_type=authorization_code&code=` + authCode + `&redirect_uri=` + conf.RedirectURL)
	request, err := http.NewRequest("POST", "https://api.reddit.com/api/v1/access_token", bytes.NewBuffer(query))
	if err != nil {
		return
	}

	//Set custom UA & Auth because Reddit is annoying
	request.Header.Add("Authorization", "Basic "+basicAuth(conf.ClientID, conf.ClientSecret))
	request.Header.Set("User-Agent", "eutenly-backend/0.1")

	//Send request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

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
	if responseJSON["error"] != nil {
		err = fmt.Errorf("reddit: %v", responseJSON["error"])
		return
	}

	accessToken = fmt.Sprintf("%v", responseJSON["access_token"])
	refreshToken = fmt.Sprintf("%v", responseJSON["refresh_token"])
	return
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
