package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var collection *mongo.Collection

func Initdb() (*mongo.Client, error) {
	// Mongodb Connection String
	clientOptions := options.Client().ApplyURI("mongodb+srv://Toby:lRgxjw2HzMEdgJAF@cluster0.jalvxqn.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")

	//Getting Username and Password from .env
	username := os.Getenv("MONGO_DB_USERNAME")
	password := os.Getenv("MONGO_DB_PASSWORD")

	// Setting auth Credentials
	clientOptions.SetAuth(options.Credential{
		Username: username,
		Password: password,
	})

	//	Connect to mongodb
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Println("connected to mongo...")

	return client, nil
}
