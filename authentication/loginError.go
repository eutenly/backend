package authentication

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func loginError(err error, service string, c echo.Context) error {
	authCookie("spotify", c)
	c.SetCookie(&http.Cookie{Name: "auth_error", Value: fmt.Sprint(err.Error()), Expires: time.Now().Add(time.Hour * 1), Path: "/"})
	return c.Redirect(302, "/connections")
}
