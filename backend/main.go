package main

import (
    "bookshop-backend/database"
    "bookshop-backend/routes"
    "fmt"
    "net/http"
    "os" // Required for os.Exit or handlers package (good practice)

    "github.com/gorilla/mux"
 "path/filepath" // Import this package
    "github.com/gorilla/handlers" // <--- 1. IMPORT THIS PACKAGE
)

func main() {
    // 1. Connect to the database
    database.Connect()

    // 2. Create the main router instance
    r := mux.NewRouter()


// 1. Define the ABSOLUTE path to the folder containing your image.
    // This MUST point directly to the static_assets folder.
    const absoluteAssetsDir = "C:/Users/User/OneDrive/Desktop/bookshop-ecommerce/static_assets"
    
    // Convert path to http.Dir format
    fsPath := http.Dir(filepath.FromSlash(absoluteAssetsDir))

    // 2. Register the static file handler using a new prefix: /static/
    // Request for: http://localhost:8081/static/go_book.jpg
    // File found at: C:/.../static_assets/go_book.jpg
    r.PathPrefix("/static/").Handler(
        http.StripPrefix("/static/", http.FileServer(fsPath)),
    )


    // 3. Register all routes onto the single router instance
    routes.BookRoutes(r)
    routes.BookIDRoutes(r)
    routes.BookCategoryRoutes(r)
    routes.Register_Customer_Routes(r)
    routes.Login_Customer_Routes(r)
    routes.CartRoutes(r) // Cart APIs
    routes.OrderRoutes(r) // Orders APIs
    routes.DiscountRoutes(r) // discount Routes
    routes.CountdownSaleRoutes(r)

    // --- 4. CORS Configuration and Middleware ---
    // In a development environment, we use "*" to allow all origins.
    // In production, change "*" to your actual frontend URL (e.g., "http://yourfrontend.com").
    allowedOrigins := handlers.AllowedOrigins([]string{"*"})
    // Allow necessary methods for a full API
    allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
    // Allow headers like Content-Type and Authorization (for login/cart)
    allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})

    // Wrap the Gorilla Mux router (r) with the CORS middleware
    handler := handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r)
    // ---------------------------------------------

    // 5. Start the server
    port := "8081"
    fmt.Printf("🚀 Server running at http://localhost:%s\n", port)
    
    // Pass the CORS-wrapped handler instead of the raw router 'r'
    err := http.ListenAndServe(":"+port, handler) 
    if err != nil {
        fmt.Println("Server failed to start:", err)
        os.Exit(1)
    }
}