package controllers

import (
    "bookshop-backend/database"
    "bookshop-backend/models"
    "encoding/json"
    "net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
    // Include all new columns in SELECT query
    rows, err := database.DB.Query(`
        SELECT 
            id, title, author, price, stock, description, image_url, 
            category, isbn, publisher, publish_year, rating, pages, language, created_at
        FROM books
    `)
    if err != nil {
        http.Error(w, "Database query error", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var books []models.Book
    for rows.Next() {
        var b models.Book
        err := rows.Scan(
            &b.ID, &b.Title, &b.Author, &b.Price, &b.Stock, &b.Description, &b.ImageURL,
            &b.Category, &b.ISBN, &b.Publisher, &b.PublishYear, &b.Rating, &b.Pages, &b.Language, &b.CreatedAt,
        )
        if err != nil {
            http.Error(w, "Error scanning database row", http.StatusInternalServerError)
            return
        }
        books = append(books, b)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

func GetBooksByCategory(w http.ResponseWriter, r *http.Request) {
    category := r.URL.Query().Get("category")
    if category == "" {
        http.Error(w, "Missing category parameter", http.StatusBadRequest)
        return
    }

    rows, err := database.DB.Query("SELECT id, title, author, price, stock, description, image_url, category FROM books WHERE category = ?", category)
    if err != nil {
        http.Error(w, "Database query error", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var books []models.Book
    for rows.Next() {
        var b models.Book
        rows.Scan(&b.ID, &b.Title, &b.Author, &b.Price, &b.Stock, &b.Description, &b.ImageURL, &b.Category)
        books = append(books, b)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}
