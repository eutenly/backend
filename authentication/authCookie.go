package authentication

import (
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func authCookie(service string, c echo.Context) {
	c.SetCookie(&http.Cookie{Name: "authed_with", Value: service, Expires: time.Now().Add(time.Hour * 1), Path: "/"})
}
