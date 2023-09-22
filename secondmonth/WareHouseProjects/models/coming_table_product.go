package models

import "time"

type CreateComingTableProduct struct {
	Category_id     int     `json:"category_id"`
	Name            string  `json:"name"`
	Price           float64 `json:"price"`
	Barcode         string  `json:"barcode"`
	Count           float64 `json:"count"`
	Coming_Table_id string  `json:"coming_table_id"`
}

type ComingTableProduct struct {
	ID              string  `json:"id"`
	Category_id     int     `json:"category_id"`
	Name            string  `json:"name"`
	Price           float64 `json:"price"`
	Barcode         string  `json:"barcode"`
	Count           float64 `json:"count"`
	TotalPrice      float64 `json:"total_price"`
	Coming_Table_id string  `json:"coming_table_id"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}

type ComingTableProductIdRequest struct {
	Id string `json:"id"`
}

type GetAllComingTableProductRequest struct {
	ComingID string    `json:"coming_id"`
	BranchID string    `json:"branch_id"`
	DateTime time.Time `json:"date_time"`
}

type GetAllComingTableProduct struct {
	ComingTableProducts []ComingTableProduct `json:"coming_table_product"`
	Count               int                  `json:"count"`
}
