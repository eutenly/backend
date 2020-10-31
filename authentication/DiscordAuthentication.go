package authentication

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

type authResp struct {
	Code  string `json:"access_token"`
	Error string `json:"error"`
}

type discordUser struct {
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	ID            string `json:"id"`
	AvatarID      string `json:"avatar"`
}

func DiscordAuthenticationRoutes(e *echo.Echo) {
	//Discord Login Route
	e.GET("/login/discord", func(c echo.Context) error {
		//Create login URL
		clientID := os.Getenv("DISCORD_CLIENT_ID")
		redirectURI := os.Getenv("WEBSERVER_URL") + "/auth/discord"
		authURL := fmt.Sprintf("https://discord.com/api/oauth2/authorize?client_id=%v&redirect_uri=%v&response_type=code&scope=identify", clientID, redirectURI)

		if c.QueryParam("redirect_to") != "" {
			sess, _ := session.Get("session", c)
			sess.Values["redirect_to"] = c.QueryParam("redirect_to")
			sess.Save(c.Request(), c.Response())
		}

		//Send user there
		return c.Redirect(http.StatusTemporaryRedirect, authURL)
	})

	//Discord Auth Route
	e.GET("/auth/discord", func(c echo.Context) error {
		//Get OAUTH2 login code
		authCode := c.QueryParam("code")

		//If no token was passed then error
		if authCode == "" {
			return c.String(http.StatusUnauthorized, "A Discord login error occured.")
		}

		//Request accessToken
		accessToken, err := authenticateDiscord(authCode)
		if err != nil {
			c.SetCookie(&http.Cookie{Name: "authed_with", Value: "discord"})
			c.SetCookie(&http.Cookie{Name: "auth_error", Value: fmt.Sprint(err.Error())})
			return c.Redirect(302, "/login-error")
		}

		//Fetch user details
		authenticatedUser, err := getDiscordUser(accessToken)
		if err != nil {
			c.SetCookie(&http.Cookie{Name: "authed_with", Value: "discord"})
			c.SetCookie(&http.Cookie{Name: "auth_error", Value: fmt.Sprint(err.Error())})
			return c.Redirect(302, "/login-error")
		}
		sess, _ := session.Get("session", c)

		//Store Discord details in Session and save
		sess.Values["authed"] = true
		sess.Values["discord_accessToken"] = accessToken
		sess.Values["discord_username"] = authenticatedUser.Username
		sess.Values["discord_discrim"] = authenticatedUser.Discriminator
		sess.Values["discord_id"] = authenticatedUser.ID
		sess.Values["discord_avatar"] = authenticatedUser.AvatarID

		//Save session
		sess.Save(c.Request(), c.Response())

		//Redirect
		if sess.Values["redirect_to"] != nil {
			return c.Redirect(302, fmt.Sprint(sess.Values["redirect_to"]))
		}

		//Respond
		return c.Redirect(302, "/connections")
	})

	//Logout
	e.GET("/logout", func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Values["authed"] = false
		sess.Save(c.Request(), c.Response())
		return c.Redirect(http.StatusTemporaryRedirect, "/?logout")
	})
}

func authenticateDiscord(authCode string) (accessToken string, rErr error) {
	//Make request to get Discord Token
	response, err := http.PostForm("https://discord.com/api/oauth2/token", url.Values{
		"client_id":     {os.Getenv("DISCORD_CLIENT_ID")},
		"client_secret": {os.Getenv("DISCORD_CLIENT_SECRET")},
		"grant_type":    {"authorization_code"},
		"code":          {authCode},
		"redirect_uri":  {os.Getenv("WEBSERVER_URL") + "/auth/discord"},
		"scope":         {"identify"}})
	if err != nil {
		return "err", err
	}
	defer response.Body.Close()

	//Format response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "err", err
	}

	//Format JSON into an object Struct
	var responseFormat authResp
	err = json.Unmarshal(body, &responseFormat)
	if err != nil {
		return "err", err
	}

	//Check for errors
	if responseFormat.Error != "" {
		return "err", fmt.Errorf(responseFormat.Error)
	}

	//Return Access Code
	return responseFormat.Code, nil
}

func getDiscordUser(accessToken string) (user discordUser, reqerr error) {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", "https://discord.com/api/users/@me", nil)
	req.Header.Add("Authorization", "Bearer "+accessToken)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		reqerr = err
		return
	}

	defer resp.Body.Close()

	//Format response body
	var userResponse discordUser
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		reqerr = err
		return
	}

	//Format JSON into an object Struct
	err = json.Unmarshal(body, &userResponse)
	if err != nil {
		reqerr = err
		return
	}

	return userResponse, nil
}
