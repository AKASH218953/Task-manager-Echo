package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Logout(c echo.Context) error {
	c.Set("session", nil)
	return c.Redirect(http.StatusFound, "/")

}
