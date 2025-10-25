package routes

import (
    "bookshop-backend/controllers"
    "github.com/gorilla/mux"
)

func BookRoutes(r *mux.Router) {
    r.HandleFunc("/api/books", controllers.GetBooks).Methods("GET")
}
