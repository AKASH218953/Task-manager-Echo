package models

type Task struct {
	ID          string `bson:"_id"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	UserID      string `bson:"user_id"`
}
