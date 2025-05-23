package database

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
	"time"
)

var (
	clientInstance *mongo.Client
	clientOnce     sync.Once
)

func GetMongoClient() *mongo.Client {
	clientOnce.Do(func() {
		// Load .env file
		if err := godotenv.Load("../.env"); err != nil {
			log.Println("No .env file found or failed to load it")
		}

		//uri := os.Getenv("MONGO_URI")
		uri := "mongodb+srv://akulasaht:tnmNEP0sO7Hxh2mD@taskcluster.qvgv6lg.mongodb.net/Todo_db"

		log.Println("mongourl", uri)
		if uri == "" {
			log.Fatal("MONGO_URI not set in environment")
		}

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		var err error
		clientInstance, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
		if err != nil {
			log.Fatalf("Mongo connection error: %v", err)
		}

		if err := clientInstance.Ping(ctx, nil); err != nil {
			log.Fatalf("Mongo ping failed: %v", err)
		}

		log.Println("MongoDB connected successfully")
	})

	return clientInstance

}
