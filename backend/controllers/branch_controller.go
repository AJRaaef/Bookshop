package controllers

import (
	"bookshop-backend/database"
	"bookshop-backend/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"log"
)

// ✅ Add a new branch
func AddBranch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var b models.Branch
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
    
    // Basic validation
    if b.Name == "" || b.BranchCode == "" {
        http.Error(w, "Name and BranchCode are required fields", http.StatusBadRequest)
        return
    }

	_, err = database.DB.Exec(
		"INSERT INTO branch (name, address, phone, branch_code) VALUES (?, ?, ?, ?)",
		b.Name, b.Address, b.Phone, b.BranchCode,
	)
	if err != nil {
		log.Println("Database error inserting branch:", err)
		http.Error(w, "Failed to insert branch", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Branch added successfully"})
}

// ✅ Get all branches
func GetAllBranches(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rows, err := database.DB.Query("SELECT branch_id, name, address, phone, branch_code FROM branch")
	if err != nil {
		log.Println("Database query error retrieving all branches:", err)
		http.Error(w, "Database query error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var branches []models.Branch
	for rows.Next() {
		var b models.Branch
		err := rows.Scan(&b.BranchID, &b.Name, &b.Address, &b.Phone, &b.BranchCode)
		if err == nil {
			branches = append(branches, b)
		} else {
            log.Println("Error scanning branch row:", err)
        }
	}

	json.NewEncoder(w).Encode(branches)
}

// ✅ Get a single branch by BranchCode
func GetBranchByCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
    // The BranchCode is typically passed as a query parameter (e.g., ?code=CBM)
    branchCode := r.URL.Query().Get("code")
    if branchCode == "" {
        http.Error(w, "Branch code parameter 'code' is required", http.StatusBadRequest)
        return
    }

	row := database.DB.QueryRow("SELECT branch_id, name, address, phone, branch_code FROM branch WHERE branch_code = ?", branchCode)

	var b models.Branch
	err := row.Scan(&b.BranchID, &b.Name, &b.Address, &b.Phone, &b.BranchCode)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Branch not found", http.StatusNotFound)
			return
		}
		log.Println("Database error retrieving branch by code:", err)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(b)
}