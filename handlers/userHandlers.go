package handlers

import (
	"encoding/json"
	"net/http"
	"shoeStore/db"
	"shoeStore/models"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is supported.", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.ID = generateID()

	_, err = db.UserCol.InsertOne(db.Ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is supported.", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required.", http.StatusBadRequest)
		return
	}

	var user models.User
	err := db.UserCol.FindOne(db.Ctx, bson.M{"id": id}).Decode(&user)
	if err != nil {
		http.Error(w, "User not found.", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
