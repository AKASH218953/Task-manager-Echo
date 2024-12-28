package controllers

import (
	"context"
	"myproject/config"
	"myproject/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func Register_user(c echo.Context) error {

	user_id := uuid.New().String()
	frist_name := c.FormValue("Frist_name")
	last_name := c.FormValue("Last_name")
	email := c.FormValue("Email")
	username := c.FormValue("Username")
	password := c.FormValue("Password")
	role := c.FormValue("Role")

	client := config.GetDBClient()
	collection := client.Database("users").Collection("users")

	filter := bson.M{"$or": []bson.M{
		{"email": email},
		{"username": username},
	}}

	var existingUser models.Users
	err := collection.FindOne(context.Background(), filter).Decode(&existingUser)
	if err == nil {

		return c.String(http.StatusBadRequest, "Email or Username already exists")
	}

	user := models.Users{
		ID:        user_id,
		FristName: frist_name,
		LastName:  last_name,
		Email:     email,
		UserName:  username,
		Password:  password,
		Role:      role,
	}

	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error saving user to database")
	}

	return c.Redirect(http.StatusAccepted, "/login")
}
