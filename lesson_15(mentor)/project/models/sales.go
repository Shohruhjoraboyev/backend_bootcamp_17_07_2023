package models

type CreateSales struct {
	Name             string
	Price            float64
	Payment_Type     int
	Status           int
	Client_id        int
	Branch_id        int
	Shop_asissent_id int
	Cashier_id       int
	Created_at       string
}

type Sales struct {
	Id               string
	Name             string
	Price            float64
	Payment_Type     int
	Status           int
	Client_id        int
	Branch_id        int
	Shop_asissent_id int
	Cashier_id       int
	Created_at       string
}

type GetAllSalesRequest struct {
	Page  int
	Limit int
	Name  string
}

type GetAllSalesResponse struct {
	Sales []Sales
	Count int
}
