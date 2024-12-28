package controllers

import (
	"log"
	"myproject/config"
	"myproject/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func Admin_dashboard(c echo.Context) error {
	client := config.GetDBClient()

	taskCollection := client.Database("tasks").Collection("tasks")

	var tasks []models.Task

	cursor, err := taskCollection.Find(c.Request().Context(), bson.M{})
	if err != nil {
		log.Println("Failed to fetch tasks:", err)
		return c.String(http.StatusInternalServerError, "Failed to fetch tasks")
	}
	defer cursor.Close(c.Request().Context())

	log.Println("Fetching tasks...")

	for cursor.Next(c.Request().Context()) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			log.Println("Failed to decode task:", err)
			return c.String(http.StatusInternalServerError, "Failed to decode task")
		}
		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Error occurred while iterating tasks:", err)
		return c.String(http.StatusInternalServerError, "Error occurred while iterating tasks")
	}

	return c.Render(http.StatusOK, "admin_dashboard.html", map[string]interface{}{
		"Tasks": tasks,
	})
}
