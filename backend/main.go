package main

import (
    "bookshop-backend/database"
    "bookshop-backend/routes"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    // 1. Connect to the database
    database.Connect()

    // 2. Create the main router instance
    r := mux.NewRouter()

    // 3. Register all routes onto the single router instance
    routes.BookRoutes(r)
    routes.BookIDRoutes(r) // Register the /api/book/{id} route
	routes.Register_Customer_Routes(r)
	routes.Login_Customer_Routes(r)


    // Cart APIs
routes.CartRoutes(r)


// Orders APIs
routes.OrderRoutes(r)


    // 4. Start the server
    port := "8081"
    fmt.Printf("🚀 Server running at http://localhost:%s\n", port)
    err := http.ListenAndServe(":"+port, r)
    if err != nil {
        fmt.Println("Server failed to start:", err)
    }
}