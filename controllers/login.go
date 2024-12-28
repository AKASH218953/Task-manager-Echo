package controllers

import (
	"context"
	"log"
	"myproject/config"
	"myproject/models"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoginHere(c echo.Context) error {

	username := c.FormValue("username")
	password := c.FormValue("password")

	client := config.GetDBClient()

	collection := client.Database("users").Collection("users")
	var user models.Users
	err := collection.FindOne(context.Background(), bson.M{
		"username": username,
		"password": password,
	}).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return c.String(http.StatusUnauthorized, "username or password are invalid")
	}
	if err != nil {
		log.Printf("error in find username and password %v ", err)
		return c.String(http.StatusInternalServerError, "Internal server error")
	}

	sess, _ := session.Get("session", c)
	sess.Values["user_id"] = user.ID
	sess.Values["role"] = user.Role
	sess.Save(c.Request(), c.Response())

	if user.Role == "Admin" {
		return c.Redirect(http.StatusSeeOther, "/admin_dashboard")
	} else if user.Role == "User" {
		return c.Redirect(http.StatusSeeOther, "/user_dashboard")
	} else {
		return c.String(http.StatusForbidden, "Invalid role")
	}

}
