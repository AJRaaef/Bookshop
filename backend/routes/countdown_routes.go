package routes

import (
	"bookshop-backend/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func CountdownSaleRoutes(r *mux.Router) {
	// GET all sales (including inactive/expired)
	r.HandleFunc("/countdown_sales", controllers.GetAllCountdownSales).Methods(http.MethodGet)
    
	// GET currently running active sale (checks time)
	r.HandleFunc("/countdown_sales/active", controllers.GetActiveCountdownSale).Methods(http.MethodGet)
    
	// POST add a new sale
	r.HandleFunc("/countdown_sales", controllers.AddCountdownSale).Methods(http.MethodPost)
    
	// PUT change the active state (requires JSON body with ID and active boolean)
	r.HandleFunc("/countdown_sales/change_active", controllers.ChangeActiveCountdownSale).Methods(http.MethodPut)
}