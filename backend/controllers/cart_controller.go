package controllers

import (
	"bookshop-backend/database"
	"bookshop-backend/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Add item to cart
func AddToCart(w http.ResponseWriter, r *http.Request) {
	var cart models.Cart
	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate input
	if cart.CustomerID == 0 || cart.BookID == 0 || cart.Quantity <= 0 {
		http.Error(w, "Customer ID, Book ID, and Quantity are required", http.StatusBadRequest)
		return
	}

	// Insert into cart
	_, err = database.DB.Exec(
		"INSERT INTO cart (customer_id, book_id, quantity, added_at) VALUES (?, ?, ?, ?)",
		cart.CustomerID, cart.BookID, cart.Quantity, time.Now(),
	)
	if err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		fmt.Println("DB error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
	fmt.Println("Added to cart:", cart)
}

// Get all cart items for a customer
func GetCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["customer_id"]

	rows, err := database.DB.Query(
		"SELECT id, customer_id, book_id, quantity, added_at FROM cart WHERE customer_id = ?",
		customerID,
	)
	if err != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var carts []models.Cart
	for rows.Next() {
		var c models.Cart
		err := rows.Scan(&c.ID, &c.CustomerID, &c.BookID, &c.Quantity, &c.AddedAt)
		if err != nil {
			http.Error(w, "Error scanning cart", http.StatusInternalServerError)
			return
		}
		carts = append(carts, c)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(carts)
}

// Remove cart item by ID
func RemoveCartItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartID := vars["id"]

	_, err := database.DB.Exec("DELETE FROM cart WHERE id = ?", cartID)
	if err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	fmt.Println("Removed cart item ID:", cartID)
}
