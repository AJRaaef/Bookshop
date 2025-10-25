package models

type Order struct {
	ID         int     `json:"id"`
	CustomerID int     `json:"customer_id"`
	BookID     int     `json:"book_id"`
	Quantity   int     `json:"quantity"`
	Total      float64 `json:"total"`
	Status     string  `json:"status"`
	OrderDate  string  `json:"order_date,omitempty"`
}
