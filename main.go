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
		http.ServeFile(w, r, "./templates/index.html")
	})

	http.HandleFunc("/loginuser", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/Login.html")
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/register.html")
	})

	http.HandleFunc("/shoes", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/shoes.html")
	})

	fs := http.FileServer(http.Dir("./templates"))
	http.Handle("/templates/", http.StripPrefix("/templates/", fs))

	log.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", nil)
}
