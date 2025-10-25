package controllers

import (
	"bookshop-backend/database"
	"bookshop-backend/models"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Register_Customer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer

	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if customer.Name == "" || customer.Email == "" || customer.Password == "" {
		http.Error(w, "Name, email and password are required", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	result, err := database.DB.Exec(
		"INSERT INTO customers (name, email, password) VALUES (?, ?, ?)",
		customer.Name, customer.Email, string(hashedPassword),
	)
	if err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		fmt.Println("DB error:", err)
		return
	}

	id, _ := result.LastInsertId()
	customer.ID = int(id)
	customer.Password = "" // hide password in response

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
	fmt.Println("Customer registered:", customer.Email)
}
