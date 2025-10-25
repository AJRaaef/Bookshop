package controllers

import (
	"bookshop-backend/database"
	"bookshop-backend/models"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Login_Customer(w http.ResponseWriter, r *http.Request) {
	var input models.Customer

	// Decode JSON body
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Get customer from DB
	var customer models.Customer
	err = database.DB.QueryRow(
		"SELECT id, name, email, password FROM customers WHERE email = ?",
		input.Email,
	).Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Password)

	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(input.Password))
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	customer.Password = "" // hide password in response

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
	fmt.Println("Customer logged in:", customer.Email)
}
