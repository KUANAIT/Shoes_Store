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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/index.html")
	})

	http.HandleFunc("/loginuser", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/Login.html")
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/register.html")
	})

	http.HandleFunc("/shoes", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/shoes.html")
	})

	fs := http.FileServer(http.Dir("./docs"))
	http.Handle("/docs/", http.StripPrefix("/docs/", fs))

	log.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", nil)
}
