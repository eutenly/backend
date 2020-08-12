package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

type connection struct {
	Service  string `json:"service"`
	Username string `json:"username"`
}

type accountDetails struct {
	Username    string       `json:"username"`
	Discrim     string       `json:"discrim"`
	ID          string       `json:"id"`
	Connections []connection `json:"connections"`
}

func MyEutenlyRoutes(e *echo.Echo) {
	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "MyEutenly API")
	})

	e.GET("/api/me", func(c echo.Context) error {
		//Get session
		sess, _ := session.Get("session", c)
		if sess.Values["authed"] != true {
			return c.String(http.StatusUnauthorized, "Not logged in")
		}

		//Create account object
		var userconnections []connection
		u := &accountDetails{
			Username:    fmt.Sprint(sess.Values["discord_username"]),
			Discrim:     fmt.Sprint(sess.Values["discord_discrim"]),
			ID:          fmt.Sprint(sess.Values["discord_id"]),
			Connections: userconnections,
		}

		//Serve
		return c.JSON(http.StatusOK, u)
	})
}
