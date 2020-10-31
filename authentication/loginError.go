package authentication

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func loginError(err error, service string, c echo.Context) error {
	c.SetCookie(&http.Cookie{Name: "authed_with", Value: "spotify"})
	c.SetCookie(&http.Cookie{Name: "auth_error", Value: fmt.Sprint(err.Error())})
	return c.Redirect(302, "/connections")
}
