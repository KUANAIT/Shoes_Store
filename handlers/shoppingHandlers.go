package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"shoeStore/db"
	"shoeStore/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateShoeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is supported.", http.StatusMethodNotAllowed)
		return
	}

	var shoe models.Shoe
	err := json.NewDecoder(r.Body).Decode(&shoe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	shoe.ID = generateID()

	_, err = db.ShoeCol.InsertOne(db.Ctx, shoe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shoe)
}

func GetAllShoes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is supported.", http.StatusMethodNotAllowed)
		return
	}

	cursor, err := db.ShoeCol.Find(db.Ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(db.Ctx)

	var shoes []models.Shoe
	for cursor.Next(db.Ctx) {
		var shoe models.Shoe
		if err := cursor.Decode(&shoe); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		shoes = append(shoes, shoe)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shoes)
}

func GetShoeByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is supported.", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required.", http.StatusBadRequest)
		return
	}

	var shoe models.Shoe
	err := db.ShoeCol.FindOne(db.Ctx, bson.M{"id": id}).Decode(&shoe)
	if err != nil {
		http.Error(w, "Shoe not found.", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shoe)
}

func DeleteShoeByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Only DELETE method is supported.", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required.", http.StatusBadRequest)
		return
	}

	_, err := db.ShoeCol.DeleteOne(db.Ctx, bson.M{"id": id})
	if err != nil {
		http.Error(w, "Shoe not found.", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateShoeByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Only PUT method is supported.", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required.", http.StatusBadRequest)
		return
	}

	var updatedShoe models.Shoe
	err := json.NewDecoder(r.Body).Decode(&updatedShoe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filter := bson.M{"id": id}
	update := bson.M{
		"$set": bson.M{
			"name":           updatedShoe.Name,
			"brand":          updatedShoe.Brand,
			"size":           updatedShoe.Size,
			"price":          updatedShoe.Price,
			"material":       updatedShoe.Material,
			"color":          updatedShoe.Color,
			"release_date":   updatedShoe.ReleaseDate,
			"discount":       updatedShoe.Discount,
			"stock_quantity": updatedShoe.StockQuantity,
			"rating":         updatedShoe.Rating,
			"sku":            updatedShoe.SKU,
			"weight":         updatedShoe.Weight,
			"features":       updatedShoe.Features,
			"category":       updatedShoe.Category,
			"shipping_info":  updatedShoe.ShippingInfo,
		},
	}

	result, err := db.ShoeCol.UpdateOne(db.Ctx, filter, update)
	if err != nil {
		http.Error(w, "Error updating shoe: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		http.Error(w, "Shoe not found.", http.StatusNotFound)
		return
	}

	var shoe models.Shoe
	err = db.ShoeCol.FindOne(db.Ctx, filter).Decode(&shoe)
	if err != nil {
		http.Error(w, "Error retrieving updated shoe: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shoe)
}

func generateID() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%d", rand.Intn(1000000))
}
