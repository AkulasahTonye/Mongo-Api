package routes

import (
	"encoding/json"
	"github.com/mongo-Api/models"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

type Handler struct {
	collection *mongo.Collection
}

var currentUser models.User

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Error Decoding JSON:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	currentUser = user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(currentUser)
}
