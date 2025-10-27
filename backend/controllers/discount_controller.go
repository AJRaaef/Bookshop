package controllers

import (
    "bookshop-backend/database"
    "bookshop-backend/models"
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

// ✅ 1. Get all discounts with book details
func GetAllDiscounts(w http.ResponseWriter, r *http.Request) {
    query := `
        SELECT d.discount_id, d.book_id, b.title, b.category, b.price, 
               d.discount_percent, d.is_active, 
               (b.price - (b.price * d.discount_percent / 100)) AS discounted_price
        FROM discount d
        JOIN books b ON d.book_id = b.id
        WHERE d.is_active = TRUE
    `
    rows, err := database.DB.Query(query)
    if err != nil {
        http.Error(w, "Database query error", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var results []map[string]interface{}
    for rows.Next() {
        var id, bookID int
        var title, category string
        var price, discountPercent, discountedPrice float64
        var isActive bool

        rows.Scan(&id, &bookID, &title, &category, &price, &discountPercent, &isActive, &discountedPrice)

        result := map[string]interface{}{
            "discount_id":      id,
            "book_id":          bookID,
            "title":            title,
            "category":         category,
            "price":            price,
            "discount_percent": discountPercent,
            "discounted_price": discountedPrice,
            "is_active":        isActive,
        }
        results = append(results, result)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(results)
}

// ✅ 2. Get discounts by category
func GetDiscountsByCategory(w http.ResponseWriter, r *http.Request) {
    category := mux.Vars(r)["category"]

    query := `
        SELECT b.id, b.title, b.category, b.price, d.discount_percent,
               (b.price - (b.price * d.discount_percent / 100)) AS discounted_price
        FROM discount d
        JOIN books b ON d.book_id = b.id
        WHERE b.category = ? AND d.is_active = TRUE
    `
    rows, err := database.DB.Query(query, category)
    if err != nil {
        http.Error(w, "Database query error", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var results []map[string]interface{}
    for rows.Next() {
        var id int
        var title, categoryName string
        var price, discountPercent, discountedPrice float64

        rows.Scan(&id, &title, &categoryName, &price, &discountPercent, &discountedPrice)

        result := map[string]interface{}{
            "book_id":          id,
            "title":            title,
            "category":         categoryName,
            "price":            price,
            "discount_percent": discountPercent,
            "discounted_price": discountedPrice,
        }
        results = append(results, result)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(results)
}

// ✅ 3. Get discount by book ID
func GetDiscountByBookID(w http.ResponseWriter, r *http.Request) {
    idStr := mux.Vars(r)["book_id"]
    bookID, _ := strconv.Atoi(idStr)

    query := `
        SELECT d.discount_id, b.title, b.price, d.discount_percent,
               (b.price - (b.price * d.discount_percent / 100)) AS discounted_price
        FROM discount d
        JOIN books b ON d.book_id = b.id
        WHERE b.id = ? AND d.is_active = TRUE
    `
    row := database.DB.QueryRow(query, bookID)

    var result map[string]interface{}
    var discountID int
    var title string
    var price, discountPercent, discountedPrice float64

    err := row.Scan(&discountID, &title, &price, &discountPercent, &discountedPrice)
    if err != nil {
        http.Error(w, "Discount not found for this book", http.StatusNotFound)
        return
    }

    result = map[string]interface{}{
        "discount_id":      discountID,
        "book_id":          bookID,
        "title":            title,
        "price":            price,
        "discount_percent": discountPercent,
        "discounted_price": discountedPrice,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}

// ✅ 4. Add a discount for an existing book
func AddDiscount(w http.ResponseWriter, r *http.Request) {
    var newDiscount models.Discount
    if err := json.NewDecoder(r.Body).Decode(&newDiscount); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Check if book exists
    var exists bool
    err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM books WHERE id = ?)", newDiscount.BookID).Scan(&exists)
    if err != nil || !exists {
        http.Error(w, "Book not found", http.StatusBadRequest)
        return
    }

    // Insert discount
    _, err = database.DB.Exec("INSERT INTO discount (book_id, discount_percent, is_active) VALUES (?, ?, ?)",
        newDiscount.BookID, newDiscount.DiscountPercent, newDiscount.IsActive)
    if err != nil {
        http.Error(w, "Failed to add discount", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "Discount added successfully"})
}


// ✅ 5. Toggle discount active/inactive
func ToggleDiscountActive(w http.ResponseWriter, r *http.Request) {
    // Get discount ID from URL
    idStr := mux.Vars(r)["discount_id"]
    discountID, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid discount ID", http.StatusBadRequest)
        return
    }

    // Update is_active to NOT current value
    res, err := database.DB.Exec(`
        UPDATE discount 
        SET is_active = NOT is_active 
        WHERE discount_id = ?`, discountID)
    if err != nil {
        http.Error(w, "Failed to update discount status", http.StatusInternalServerError)
        return
    }

    rowsAffected, _ := res.RowsAffected()
    if rowsAffected == 0 {
        http.Error(w, "Discount not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Discount active status toggled successfully",
    })
}
