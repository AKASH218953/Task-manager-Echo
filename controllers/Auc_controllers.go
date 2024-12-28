package controllers

import "github.com/labstack/echo/v4"

func Homepage(c echo.Context) error {

	return c.File("templates/Homepage.html")
}
func Register(c echo.Context) error {

	return c.File("templates/Register.html")
}
func Login(c echo.Context) error {
	return c.File("templates/login.html")

}
