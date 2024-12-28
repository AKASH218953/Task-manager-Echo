package controllers

import (
	"context"
	"log"
	"myproject/config"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteTask(c echo.Context) error {
	taskID := c.Param("id") // Get task ID from URL

	client := config.GetDBClient()
	collection := client.Database("tasks").Collection("tasks")

	filter := bson.M{"_id": taskID}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Printf("Error deleting task: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to delete task")
	}

	return c.Redirect(http.StatusSeeOther, "/user_dashboard")
}
