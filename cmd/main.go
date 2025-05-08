package main

import (
	"context"
	"github.com/mongo-Api/database"
	"github.com/mongo-Api/routes"
	"github.com/mongo-Api/services"
	"log"
	"net/http"
	"time"
)

type Application struct {
	Models services.Models
}

func main() {
	mongoClient, err := database.Initdb()
	if err != nil {
		log.Panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	services.New(mongoClient)
	log.Println("Server running in port:", 8080)
	log.Fatal(http.ListenAndServe(":8080", routes.CreateRouter()))
}
