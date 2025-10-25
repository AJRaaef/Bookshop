package routes

import (
	"bookshop-backend/controllers"
	"github.com/gorilla/mux"
)

func OrderRoutes(r *mux.Router) {
	r.HandleFunc("/api/orders", controllers.PlaceOrder).Methods("POST")
	r.HandleFunc("/api/orders/{customer_id}", controllers.GetOrders).Methods("GET")
	r.HandleFunc("/api/orders/status/{id}", controllers.UpdateOrderStatus).Methods("PUT") // optional admin
}
