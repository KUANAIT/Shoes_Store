package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client   *mongo.Client
	ShoeCol  *mongo.Collection
	UserCol  *mongo.Collection
	OrderCol *mongo.Collection
	Ctx      = context.Background()
)

func Init() {
	var err error
	Client, err = mongo.Connect(Ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ShoeCol = Client.Database("shoeStore").Collection("shoes")
	UserCol = Client.Database("shoeStore").Collection("users")
	OrderCol = Client.Database("shoeStore").Collection("orders")
}
