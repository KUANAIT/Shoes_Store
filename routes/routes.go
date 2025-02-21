package routes

import (
	"net/http"
	"shoeStore/handlers"
)

func SetupRoutes() {
	http.HandleFunc("/user/create", handlers.CreateUserHandler)
	http.HandleFunc("/user/getbyid", handlers.GetUserByID)
	http.HandleFunc("/order/create", handlers.CreateOrderHandler)
	http.HandleFunc("/order/getbyuser", handlers.GetOrdersByUserID)

	http.HandleFunc("/create", handlers.CreateShoeHandler)
	http.HandleFunc("/getall", handlers.GetAllShoes)
	http.HandleFunc("/getbyid", handlers.GetShoeByID)
	http.HandleFunc("/delete", handlers.DeleteShoeByID)
	http.HandleFunc("/update", handlers.UpdateShoeByID)
}
