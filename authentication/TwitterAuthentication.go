package authentication

import (
	"fmt"
	"net/http"
	"os"

	twitter "github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	twitterAuth "github.com/dghubble/oauth1/twitter"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

func TwitterAuthenticationRoutes(e *echo.Echo) {
	config := oauth1.Config{
		ConsumerKey:    os.Getenv("TWITTER_KEY"),
		ConsumerSecret: os.Getenv("TWITTER_SECRET"),
		CallbackURL:    os.Getenv("WEBSERVER_URL") + "/auth/twitter",
		Endpoint:       twitterAuth.AuthorizeEndpoint,
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

		//Login as user to get Twitter ID
		config := oauth1.NewConfig(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"))
		token := oauth1.NewToken(accessToken, accessSecret)
		httpClient := config.Client(oauth1.NoContext, token)

		// Twitter client
		twitterClient := twitter.NewClient(httpClient)

		authedUser, _, err := twitterClient.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{})
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		//database.FindByID(fmt.Sprintf("%v", sess.Values["discord_id"]))

		return c.String(http.StatusAccepted, fmt.Sprintf("UserID: %v; accessToken: %v; accessSecret %v;", authedUser.ID, accessToken, accessSecret))
	})

}
