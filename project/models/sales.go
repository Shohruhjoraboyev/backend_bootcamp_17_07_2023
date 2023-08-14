package models

type CreateSales struct {
	ClientName      string
	BranchId        int
	Shop_asissentId int
	CashierId       int
	Price           float64
	PaymentType     Payment
	Status          Status
	CreatedAt       string
}

type Status string

const (
	Success Status = "success"
	Cancel  Status = "cancel"
)

type Payment string

const (
	Card Payment = "card"
	Cash Payment = "cash"
)

type Sales struct {
	Id              string
	ClientName      string
	BranchId        int
	Shop_asissentId int
	CashierId       int
	Price           float64
	PaymentType     Payment
	Status          Status
	CreatedAt       string
}

type GetAllSalesRequest struct {
	Page       int
	Limit      int
	ClientName string
}

type GetAllSalesResponse struct {
	Sales []Sales
	Count int
}
