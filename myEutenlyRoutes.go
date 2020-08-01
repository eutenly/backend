package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

type myAccount struct {
	Username    string       `json:"discord_username"`
	ID          string       `json:"discord_id"`
	Connections []connection `json:"connections"`
}

type connection struct {
	Service  string `json:"service"`
	Username string `json:"username"`
}

func myEutenlyRoutes(e *echo.Echo) {
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
		u := &myAccount{
			Username:    fmt.Sprint(sess.Values["discord_username"]),
			ID:          fmt.Sprint(sess.Values["discord_id"]),
			Connections: userconnections,
		}

		//Serve
		return c.JSON(http.StatusOK, u)
	})
	// e.GET("/api/demo-login", func(c echo.Context) error {
	// 	sess, _ := session.Get("session", c)
	// 	sess.Values["authed"] = true
	// 	sess.Values["discord_username"] = "Maia#1234"
	// 	sess.Values["discord_id"] = "149862827027464193"
	// 	sess.Save(c.Request(), c.Response())
	// 	return c.String(http.StatusAccepted, "Logged in")
	// })
}
