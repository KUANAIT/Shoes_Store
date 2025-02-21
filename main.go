package main

import (
	"log"
	"net/http"

	"shoeStore/db"
	"shoeStore/routes"
)

func main() {
	db.Init()
	routes.SetupRoutes()

	log.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", nil)
}
