package routes

import (
	"bookshop-backend/controllers"

	"github.com/gorilla/mux"
)

func CartRoutes(r *mux.Router) {
	r.HandleFunc("/api/cart", controllers.AddToCart).Methods("POST")
	r.HandleFunc("/api/cart/{customer_id}", controllers.GetCart).Methods("GET")
	r.HandleFunc("/api/cart/remove/{id}", controllers.RemoveCartItem).Methods("DELETE")
}
