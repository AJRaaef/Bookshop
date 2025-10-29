package models

import "time"

type CountdownSale struct {
    CountdownSaleID    int       `json:"countdown_sale_id"`
    Name               string    `json:"name"`
    DiscountPercentage float64   `json:"discount_percentage"`
    StartTime          time.Time `json:"start_time"`
    EndTime            time.Time `json:"end_time"`
    Active             bool      `json:"active"`
}