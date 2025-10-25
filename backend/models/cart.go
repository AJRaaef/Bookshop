package models

type Cart struct {
	ID         int    `json:"id"`           // matches `id` column
	CustomerID int    `json:"customer_id"`  // matches `customer_id` column
	BookID     int    `json:"book_id"`      // matches `book_id` column
	Quantity   int    `json:"quantity"`     // matches `quantity` column
	AddedAt    string `json:"added_at,omitempty"` // matches `added_at` column
}
