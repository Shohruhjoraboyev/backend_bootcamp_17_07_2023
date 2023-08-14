package models

type CreateStaffRequest struct {
	BranchId int
	TariffId int
	Type     StaffType
	Name     string
	Balance  float64
}

type Staff struct {
	Id       string
	BranchId int
	TariffId int
	Type     string
	Name     string
	Balance  float64
}
type StaffType string

const (
	Cashier       StaffType = "cashier"
	ShopAssistant StaffType = "shop_assistant"
)

type GetStaffByIdRequest struct {
	Id string
}

type GetAllStaffRequest struct {
	BranchId    int
	TariffId    int
	Type        StaffType
	Name        string
	BalanceFrom float64
	BalanceTo   float64
	Page        int
	Limit       int
}

type GetAllStaffResponse struct {
	Staffs []Staff
	Count  int
}
