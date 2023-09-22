package models

type CreateRemain struct {
	Branch_id  int     `json:"branch_id"`
	Cagory_id  int     `json:"category_id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Barcode    string  `json:"barcode"`
	Count      float64 `json:"count"`
	TotalPrice float64 `json:"total_price"`
}

type Remain struct {
	ID         string  `json:"id"`
	Branch_id  int     `json:"branch_id"`
	Cagory_id  int     `json:"category_id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Barcode    string  `json:"barcode"`
	Count      float64 `json:"count"`
	TotalPrice float64 `json:"total_price"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

type RemainIdRequest struct {
	Id string `json:"id"`
}

type GetAllRemainRequest struct {
	Name    string  `json:"name"`
	Barcode string  `json:"barcode"`
	Count   float64 `json:"count"`
}

type GetAllRemain struct {
	Remainings []Remain `json:"remaining"`
	Count      int      `json:"count"`
}
