package controllers

import (
    "bookshop-backend/database"
    "encoding/json"
    "net/http"
)

// GetCategories returns all unique book categories
func GetCategories(w http.ResponseWriter, r *http.Request) {
    rows, err := database.DB.Query("SELECT DISTINCT category FROM books WHERE category IS NOT NULL AND category != ''")
    if err != nil {
        http.Error(w, "Error fetching categories", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var categories []string
    for rows.Next() {
        var category string
        rows.Scan(&category)
        categories = append(categories, category)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(categories)
}
