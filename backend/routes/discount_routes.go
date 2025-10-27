package routes

import (
    "bookshop-backend/controllers"
    "github.com/gorilla/mux"
)

func DiscountRoutes(r *mux.Router) {
    r.HandleFunc("/api/discounts", controllers.GetAllDiscounts).Methods("GET")
    r.HandleFunc("/api/discounts/category/{category}", controllers.GetDiscountsByCategory).Methods("GET")
    r.HandleFunc("/api/discounts/book/{book_id}", controllers.GetDiscountByBookID).Methods("GET")
    r.HandleFunc("/api/discounts", controllers.AddDiscount).Methods("POST")
	r.HandleFunc("/api/discounts/toggle/{discount_id}", controllers.ToggleDiscountActive).Methods("PUT")

}
