package routes

import (
	"bookshop-backend/controllers"
	"github.com/gorilla/mux"
)

func Register_Customer_Routes(r *mux.Router) {
	r.HandleFunc("/api/customers", controllers.Register_Customer).Methods("POST")
}
