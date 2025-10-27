package routes

import(
	"bookshop-backend/controllers"

	"github.com/gorilla/mux"
)


func BookCategoryRoutes(r *mux.Router) {
    // Defines the route with the variable {id}
    r.HandleFunc("/api/categories", controllers.GetCategories).Methods("GET")
}

