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

// Place an order
func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate
	if order.CustomerID == 0 || order.BookID == 0 || order.Quantity <= 0 || order.Total <= 0 {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	// Default status
	order.Status = "pending"
	order.OrderDate = time.Now().Format("2006-01-02 15:04:05")

	// Insert into DB
	_, err = database.DB.Exec(
		"INSERT INTO orders (customer_id, book_id, quantity, total, status, order_date) VALUES (?, ?, ?, ?, ?, ?)",
		order.CustomerID, order.BookID, order.Quantity, order.Total, order.Status, order.OrderDate,
	)
	if err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		fmt.Println("DB error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
	fmt.Println("Order placed:", order)
}

// Get all orders for a customer
func GetOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["customer_id"]

	rows, err := database.DB.Query(
		"SELECT id, customer_id, book_id, quantity, total, status, order_date FROM orders WHERE customer_id = ?",
		customerID,
	)
	if err != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var o models.Order
		err := rows.Scan(&o.ID, &o.CustomerID, &o.BookID, &o.Quantity, &o.Total, &o.Status, &o.OrderDate)
		if err != nil {
			http.Error(w, "Error scanning order", http.StatusInternalServerError)
			return
		}
		orders = append(orders, o)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

// Optional: Update order status (admin)
func UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["id"]

	var status struct {
		Status string `json:"status"`
	}
	err := json.NewDecoder(r.Body).Decode(&status)
	if err != nil || status.Status == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec(
		"UPDATE orders SET status = ? WHERE id = ?",
		status.Status, orderID,
	)
	if err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println("Order status updated ID:", orderID, "to", status.Status)
}
