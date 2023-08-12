package models

type BranchTransactionCount struct {
	BranchID   int
	BranchName string
	Count      int
}

type BranchProduct struct {
	BranchId  int `json:"branch_id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
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
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Product struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryID int    `json:"category_id"`
}
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
