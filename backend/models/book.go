package models

type Book struct {
    ID           int     `json:"id"`
    Title        string  `json:"title"`
    Author       string  `json:"author"`
    Price        float64 `json:"price"`
    Stock        int     `json:"stock"`
    Description  string  `json:"description"`
    ImageURL     string  `json:"image_url"`
    Category     string  `json:"category"`
    ISBN         string  `json:"isbn"`
    Publisher    string  `json:"publisher"`
    PublishYear  int     `json:"publish_year"`
    Rating       float64 `json:"rating"`
    Pages        int     `json:"pages"`
    Language     string  `json:"language"`
    CreatedAt    string  `json:"created_at"`
}
