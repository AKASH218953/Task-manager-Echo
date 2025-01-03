package middleware

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func LoginMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		sess, err := session.Get("session", c)
		if err != nil || sess.Values["user_id"] == nil {

			return c.Redirect(http.StatusFound, "/login")
		}

		return next(c)
	}
}
