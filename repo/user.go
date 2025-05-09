package repo

import (
	"context"
	"github.com/mongo-Api/database"
	"github.com/mongo-Api/models"
	"github.com/mongo-Api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

func Save(user models.User) (*mongo.InsertOneResult, error) {
	// Hash the password before saving
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return nil, err
	}
	user.Password = hashedPassword

	client := database.GetMongoClient()

	// Use a 15-second timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("Todo_db").Collection("USER")

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return nil, err
	}

	return result, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	client := database.GetMongoClient()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("Todo_db").Collection("USER")

	var user models.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetAllUsers retrieves all users from the database

func GetAllUsers() ([]models.User, error) {
	client := database.GetMongoClient()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("Todo_db").Collection("USER")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

// GetUserByID retrieves a single user by ID
func GetUserByID(id primitive.ObjectID) (*models.User, error) {
	client := database.GetMongoClient()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("Todo_db").Collection("USER")

	var user models.User
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser updates an existing user
func UpdateUser(user models.User) error {
	client := database.GetMongoClient()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("Todo_db").Collection("USER")

	update := bson.M{
		"$set": bson.M{
			"name":     user.Name,
			"email":    user.Email,
			"age":      user.Age,
			"password": user.Password,
		},
	}

	_, err := collection.UpdateOne(ctx, bson.M{"_id": user.ID}, update)
	return err
}

// DeleteUser deletes a user by ID
func DeleteUser(id primitive.ObjectID) error {
	client := database.GetMongoClient()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("Todo_db").Collection("USER")

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
