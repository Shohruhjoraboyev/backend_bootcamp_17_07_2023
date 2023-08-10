package models

type BranchTransactionCount struct {
	BranchID   int
	BranchName string
	Count      int
}
type Transaction struct {
	ID        int    `json:"id"`
	BranchID  int    `json:"branch_id"`
	ProductID int    `json:"product_id"`
	Type      string `json:"type"`
	Quantity  int    `json:"quantity"`
	CreatedAt string `json:"created_at"`
}
type Branch struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Adress string `json:"adress"`
}

type Products struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryId int    `json:"category_id"`
}

type BranchProductPrice struct {
	BranchID   int
	BranchName string
	Sum        int
}
