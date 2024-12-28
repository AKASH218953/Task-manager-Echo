package controllers

import (
	"context"
	"log"
	"myproject/config"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateTask(c echo.Context) error {
	taskID := c.Param("id")
	taskName := c.FormValue("taskname")
	taskDescription := c.FormValue("taskdescription")

	client := config.GetDBClient()
	collection := client.Database("tasks").Collection("tasks")

	filter := bson.M{"_id": taskID}
	update := bson.M{
		"$set": bson.M{
			"name":        taskName,
			"description": taskDescription,
		},
	}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("Error updating task: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to update task")
	}

	return c.Redirect(http.StatusSeeOther, "/user_dashboard")
}
