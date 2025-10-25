package routes

import (
    "bookshop-backend/controllers"
    "github.com/gorilla/mux"
)

func BookIDRoutes(r *mux.Router) {
    // Defines the route with the variable {id}
    r.HandleFunc("/api/book/{id}", controllers.GetBookByID).Methods("GET")
}