package middleware

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func UserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		sess, err := session.Get("session", c)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Unable to fetch session")
		}

		role, ok := sess.Values["role"].(string)
		if !ok || role != "User" {
			return c.String(http.StatusForbidden, "Access denied: Users only")
		}
		return next(c)
	}
}
