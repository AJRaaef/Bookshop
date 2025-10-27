package models

type Discount struct {
    DiscountID      int     `json:"discount_id"`
    BookID          int     `json:"book_id"`
    DiscountPercent float64 `json:"discount_percent"`
    IsActive        bool    `json:"is_active"`
    CreatedAt       string  `json:"created_at"`
}
