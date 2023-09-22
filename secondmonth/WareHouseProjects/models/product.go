package models

type CreateProduct struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Barcode     string  `json:"barcode"`
	Category_id int     `json:"category_id"`
}

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Barcode     string  `json:"barcode"`
	Category_id int     `json:"category_id"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type ProductIdRequest struct {
	Id string `json:"id"`
}

type GetAllProductRequest struct {
	Name    string `json:"name"`
	Barcode string `json:"barcode"`
}

type GetAllProduct struct {
	Products []Product `json:"product"`
	Count    int       `json:"count"`
}
