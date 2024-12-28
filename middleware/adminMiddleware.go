package middleware

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		sess, err := session.Get("session", c)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Session error")
		}

		role, ok := sess.Values["role"].(string)
		if !ok || role != "Admin" {
			return c.String(http.StatusForbidden, "Access denied: Admins only")
		}
		return next(c)
	}
}
