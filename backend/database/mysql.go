package database

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
    dsn := "root:@tcp(127.0.0.1:3306)/bookshopdb"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal("Error connecting to database:", err)
    }

    if err := db.Ping(); err != nil {
        log.Fatal("Database unreachable:", err)
    }

    fmt.Println("Connected to MySQL successfully ✅")
    DB = db
}
