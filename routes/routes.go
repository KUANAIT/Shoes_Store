package routes

import (
	"net/http"
	"shoeStore/handlers"
)

func SetupRoutes() {
	http.HandleFunc("/login", handlers.LoginHandler)

	http.HandleFunc("/user/create", handlers.CreateUserHandler)
	http.HandleFunc("/user/getbyid", handlers.GetUserByID)
	http.HandleFunc("/user/update", handlers.UpdateUserHandler)
	http.HandleFunc("/user/delete", handlers.DeleteUserHandler)

	http.HandleFunc("/order/create", handlers.CreateOrderHandler)
	http.HandleFunc("/order/getbyuser", handlers.GetOrdersByUserID)
	http.HandleFunc("/order/createforuser", handlers.CreateOrderForUserByUsername)

	http.HandleFunc("/create", handlers.CreateShoeHandler)
	http.HandleFunc("/getall", handlers.GetAllShoes)
	http.HandleFunc("/getbyid", handlers.GetShoeByID)
	http.HandleFunc("/delete", handlers.DeleteShoeByID)
	http.HandleFunc("/update", handlers.UpdateShoeByID)
}
