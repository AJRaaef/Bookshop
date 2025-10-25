package routes

import (
	"bookshop-backend/controllers"
	"github.com/gorilla/mux"
)

func Login_Customer_Routes(r *mux.Router) {
	r.HandleFunc("/api/login", controllers.Login_Customer).Methods("POST")
}
