package services

import "go.mongodb.org/mongo-driver/mongo"

type Task struct {
	ID     string `json:"id,omitempty" bson:"id,omitempty"`
	Name   string `json:"name,omitempty" bson:"name,omitempty"`
	Email  string `json:"email,omitempty" bson:"email,omitempty"`
	Age    int    `json:"age,omitempty" bson:"age,omitempty"`
	UserID int    `json:"user_id,omitempty" bson:"user_id,omitempty"`
}

var client *mongo.Client

func New(mongo *mongo.Client) Task {
	client = mongo

	return Task{}
}
