package controllers

import (
    "bookshop-backend/database"
    "bookshop-backend/models"
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

func GetBookByID(w http.ResponseWriter, r *http.Request) {
    // 1. Get the ID from the URL path variables
    vars := mux.Vars(r)
    idStr := vars["id"]

    // 2. Validate and convert the ID to an integer
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid book ID. Must be a number.", http.StatusBadRequest)
        fmt.Println("Invalid book ID:", idStr)
        return
    }

    fmt.Println("GetBookByID called with ID:", id)

    var book models.Book

    // 3. Query the database for the single book
    err = database.DB.QueryRow(
        "SELECT id, title, author, price, stock, description, image_url FROM books WHERE id=?",
        id,
    ).Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Stock, &book.Description, &book.ImageURL)

    // 4. Handle database query results/errors
    if err != nil {
        if err.Error() == "sql: no rows in result set" {
            // Book not found
            http.Error(w, "Book not found", http.StatusNotFound)
            fmt.Println("Book not found with ID:", id)
        } else {
            // General database error
            http.Error(w, "Database query error", http.StatusInternalServerError)
            fmt.Println("Query error:", err)
        }
        return
    }

    // 5. Success: return the book as JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(book)
    fmt.Println("Book fetched successfully:", book.Title)
}