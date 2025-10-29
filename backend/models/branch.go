package models

type Branch struct {
    BranchID   int    `json:"branch_id"`
    Name       string `json:"name"`
    Address    string `json:"address"`
    Phone      string `json:"phone"`
    BranchCode string `json:"branch_code"`
}