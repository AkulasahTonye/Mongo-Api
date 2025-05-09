package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	Name     string             `json:"name" `
	Email    string             `json:"email" `
	Age      int                `json:"age"`
	Password string             `json:"password"`
}
