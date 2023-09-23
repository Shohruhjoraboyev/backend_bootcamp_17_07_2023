package models

type CreateComingTableProduct struct {
	Category_id     string  `json:"category_id"`
	Name            string  `json:"name"`
	Price           float64 `json:"price"`
	Barcode         string  `json:"barcode"`
	Count           float64 `json:"count"`
	Coming_Table_id string  `json:"coming_table_id"`
}

type ComingTableProduct struct {
	ID              string  `json:"id"`
	Category_id     string  `json:"category_id"`
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

type UpdateComingTableProduct struct {
	ID              string  `json:"id"`
	Category_id     string  `json:"category_id"`
	Name            string  `json:"name"`
	Price           float64 `json:"price"`
	Barcode         string  `json:"barcode"`
	Count           float64 `json:"count"`
	TotalPrice      float64 `json:"total_price"`
	Coming_Table_id string  `json:"coming_table_id"`
}

type GetAllComingTableProductRequest struct {
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetAllComingTableProductResponse struct {
	ComingTableProducts []ComingTableProduct `json:"coming_table_product"`
	Count               int                  `json:"count"`
}
