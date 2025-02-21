package handlers

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"shoeStore/db"
	"shoeStore/models"
)

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is supported.", http.StatusMethodNotAllowed)
		return
	}

	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	order.ID = generateID()

	_, err = db.OrderCol.InsertOne(db.Ctx, order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filter := bson.M{"id": order.UserID}
	update := bson.M{"$push": bson.M{"order_ids": order.ID}}

	_, err = db.UserCol.UpdateOne(db.Ctx, filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

func GetOrdersByUserID(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "User ID is required.", http.StatusBadRequest)
		return
	}

	cursor, err := db.OrderCol.Find(db.Ctx, bson.M{"user_id": userID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(db.Ctx)

	var orders []models.Order
	for cursor.Next(db.Ctx) {
		var order models.Order
		if err := cursor.Decode(&order); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		orders = append(orders, order)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func CreateOrderForUserByUsername(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is supported.", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		Username string       `json:"username"`
		Order    models.Order `json:"order"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	err = db.UserCol.FindOne(db.Ctx, bson.M{"username": request.Username}).Decode(&user)
	if err != nil {
		http.Error(w, "User not found.", http.StatusNotFound)
		return
	}

	request.Order.UserID = user.ID
	request.Order.ID = generateID()

	_, err = db.OrderCol.InsertOne(db.Ctx, request.Order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filter := bson.M{"id": user.ID}
	update := bson.M{"$push": bson.M{"order_ids": request.Order.ID}}

	_, err = db.UserCol.UpdateOne(db.Ctx, filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(request.Order)
}
