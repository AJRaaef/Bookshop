package controllers

import (
	"bookshop-backend/database"
	"bookshop-backend/models"
	"database/sql"
	"encoding/json"
	"net/http"
"log"
    "time"
)

// ✅ Get all countdown sales (UPDATED FOR TIME SCANNING)
func GetAllCountdownSales(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    rows, err := database.DB.Query("SELECT countdown_sale_id, name, discount_percentage, start_time, end_time, active FROM countdown_sale")
    if err != nil {
        log.Println("Database query error in GetAllCountdownSales:", err)
        http.Error(w, "Database query error", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var sales []models.CountdownSale
    // Define the time layout matching your database format: YYYY-MM-DD HH:MM:SS
    const layout = "2006-01-02 15:04:05"

    for rows.Next() {
        var s models.CountdownSale
        var startStr, endStr string // Temporary strings to receive time data

        // *** CRITICAL CHANGE: Scan start_time and end_time into strings ***
        err := rows.Scan(
            &s.CountdownSaleID, 
            &s.Name, 
            &s.DiscountPercentage, 
            &startStr, // <-- Scan into string
            &endStr,   // <-- Scan into string
            &s.Active,
        )
        
        if err != nil {
            log.Println("Error scanning sale row:", err)
            continue // Skip this row and try the next one
        }

        // *** CONVERSION STEP: Manually parse strings to time.Time ***
        s.StartTime, err = time.Parse(layout, startStr)
        if err != nil {
            log.Println("Time Parsing Error (StartTime) in GetAll:", err)
            // Decide how to handle this error: skip the row or return an error
            continue 
        }
        
        s.EndTime, err = time.Parse(layout, endStr)
        if err != nil {
            log.Println("Time Parsing Error (EndTime) in GetAll:", err)
            continue
        }

        // Only append the sale if both scans and parses were successful (now handled by 'continue')
        sales = append(sales, s)
    }
    
    // Check for errors encountered during iteration
    if err := rows.Err(); err != nil {
        log.Println("Error after row iteration:", err)
        // You may choose to proceed or send a 500 error here
    }

    json.NewEncoder(w).Encode(sales)
}
// ✅ Get currently active and running sale
func GetActiveCountdownSale(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    now := time.Now()
    
    // We will use string placeholders for scanning the time columns temporarily
    var startStr, endStr string 
    var s models.CountdownSale
    
    // Query remains the same, using time.Time placeholder values
    query := `
        SELECT countdown_sale_id, name, discount_percentage, start_time, end_time, active 
        FROM countdown_sale 
        WHERE active = 1
          AND start_time <= ? 
          AND end_time >= ?
        LIMIT 1
    `
    // Pass time.Time objects to the driver
    row := database.DB.QueryRow(query, now, now)

    // *** CRITICAL CHANGE HERE: Scan start_time and end_time into strings ***
    err := row.Scan(
        &s.CountdownSaleID, 
        &s.Name, 
        &s.DiscountPercentage, 
        &startStr, // <--- Scan into string
        &endStr,   // <--- Scan into string
        &s.Active,
    )

    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "No active sale found currently running", http.StatusNotFound)
            return
        }
        log.Println("SQL Scan Error:", err) // Keep the log for other errors
        http.Error(w, "Database error retrieving active sale (Scan failed)", http.StatusInternalServerError)
        return
    }

    // *** CONVERSION STEP: Manually parse strings to time.Time ***
    // NOTE: The layout "2006-01-02 15:04:05" matches your sample data format (YYYY-MM-DD HH:MM:SS)
    layout := "2006-01-02 15:04:05"
    
    s.StartTime, err = time.Parse(layout, startStr)
    if err != nil {
        log.Println("Time Parsing Error (StartTime):", err)
        http.Error(w, "Internal server error: Failed to parse start time", http.StatusInternalServerError)
        return
    }
    
    s.EndTime, err = time.Parse(layout, endStr)
    if err != nil {
        log.Println("Time Parsing Error (EndTime):", err)
        http.Error(w, "Internal server error: Failed to parse end time", http.StatusInternalServerError)
        return
    }

    // The 's' struct now has valid time.Time fields
    json.NewEncoder(w).Encode(s)
}

// ✅ Add new countdown sale
func AddCountdownSale(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s models.CountdownSale
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, "Invalid input: Could not decode request body", http.StatusBadRequest)
		return
	}

	// Basic Validation: Ensure end time is after start time
	if s.EndTime.Before(s.StartTime) {
		http.Error(w, "Invalid input: End time must be after start time", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec(
		"INSERT INTO countdown_sale (name, discount_percentage, start_time, end_time, active) VALUES (?, ?, ?, ?, ?)",
		s.Name, s.DiscountPercentage, s.StartTime, s.EndTime, s.Active,
	)
	if err != nil {
		http.Error(w, "Failed to insert sale into database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Countdown sale added successfully"})
}

// ✅ Change active state (activate/deactivate sale)
func ChangeActiveCountdownSale(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type Request struct {
		CountdownSaleID int  `json:"countdown_sale_id"`
		Active          bool `json:"active"`
	}

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
    
    if req.CountdownSaleID <= 0 {
        http.Error(w, "Invalid countdown_sale_id provided", http.StatusBadRequest)
        return
    }

	_, err = database.DB.Exec("UPDATE countdown_sale SET active = ? WHERE countdown_sale_id = ?", req.Active, req.CountdownSaleID)
	if err != nil {
		http.Error(w, "Failed to update active state in database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Sale active state updated successfully"})
}