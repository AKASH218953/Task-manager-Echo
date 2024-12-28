package controllers

import (
	"context"
	"myproject/config"
	"myproject/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func CreateTask(c echo.Context) error {
	sess, _ := session.Get("session", c)
	userID, ok := sess.Values["user_id"].(string)
	if !ok {
		return c.String(http.StatusForbidden, "User not logged in")
	}

	task := models.Task{
		ID:          uuid.New().String(),
		Name:        c.FormValue("taskname"),
		Description: c.FormValue("taskdescription"),
		UserID:      userID,
	}

	client := config.GetDBClient()
	collection := client.Database("tasks").Collection("tasks")
	_, err := collection.InsertOne(context.Background(), task)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to create task")
	}

	return c.Redirect(http.StatusSeeOther, "/user_dashboard")
}
