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
)

func User_dashboard(c echo.Context) error {

	sess, _ := session.Get("session", c)
	userID, ok := sess.Values["user_id"].(string)
	if !ok {
		return c.String(http.StatusForbidden, "User not logged in")
	}

	client := config.GetDBClient()
	collection := client.Database("tasks").Collection("tasks")

	var tasks []models.Task
	cursor, err := collection.Find(context.Background(), bson.M{"user_id": userID})
	if err != nil {
		log.Printf("Error querying tasks: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to fetch tasks")
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var task models.Task
		if err := cursor.Decode(&task); err == nil {
			tasks = append(tasks, task)
		}
	}

	return c.Render(http.StatusOK, "user_dashboard.html", map[string]interface{}{
		"Tasks": tasks,
	})
}
