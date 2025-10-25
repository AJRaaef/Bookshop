package models

type Book struct {
    ID          int     `json:"id"`
    Title       string  `json:"title"`
    Author      string  `json:"author"`
    Price       float64 `json:"price"`
    Stock       int     `json:"stock"`
    Description string  `json:"description"`
    ImageURL    string  `json:"image_url"`
}
