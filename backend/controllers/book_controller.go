package controllers

import (
    "bookshop-backend/database"
    "bookshop-backend/models"
    "encoding/json"
    "net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
    rows, err := database.DB.Query("SELECT id, title, author, price, stock, description, image_url FROM books")
    if err != nil {
        http.Error(w, "Database query error", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var books []models.Book
    for rows.Next() {
        var b models.Book
        rows.Scan(&b.ID, &b.Title, &b.Author, &b.Price, &b.Stock, &b.Description, &b.ImageURL)
        books = append(books, b)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}
