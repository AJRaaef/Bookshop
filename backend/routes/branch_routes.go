package routes

import (
	"bookshop-backend/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func BranchRoutes(r *mux.Router) {
	// GET all branches
	r.HandleFunc("/branches", controllers.GetAllBranches).Methods(http.MethodGet)
    
	// GET branch by code (e.g., /branches/code?code=CBM)
	r.HandleFunc("/branches/code", controllers.GetBranchByCode).Methods(http.MethodGet)
    
	// POST add a new branch
	r.HandleFunc("/branches", controllers.AddBranch).Methods(http.MethodPost)
}