package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectDB() {

	// mongoURI := os.Getenv("MONGO_URI")
	// log.Println("MongoDB URI:", os.Getenv("MONGO_URL"))
	URL := "mongodb+srv://Akash:Akash@taskmanger.ajbin.mongodb.net/?retryWrites=true&w=majority&appName=taskmanger"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(URL))
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB!")
}

func GetDBClient() *mongo.Client {
	return client
}
