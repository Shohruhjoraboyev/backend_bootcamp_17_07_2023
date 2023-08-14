package storage

import "project/models"

type StorageI interface {
	Branch() BranchesI
	Staff() StaffI
	Sales() SalesI
	Transaction() TransactionI
	StaffTarif() StaffTarifI
}

type BranchesI interface {
	CreateBranch(models.CreateBranch) (string, error)
	UpdateBranch(models.Branch) (string, error)
	GetBranch(models.IdRequest) (models.Branch, error)
	GetAllBranch(models.GetAllBranchRequest) (models.GetAllBranch, error)
	DeleteBranch(models.IdRequest) (string, error)
}
type StaffTarifI interface {
	CreateStaffTarif(models.CreateStaffTarifRequest) (string, error)
	UpdateStaffTarif(req models.StaffTarif) (string, error)
	GetStaffTarif(req models.IdRequest) (resp models.StaffTarif, err error)
	GetAllStaffTarif(req models.GetAllStaffTarifRequest) (resp models.GetAllStaffTarif, err error)
	DeleteStaffTarif(req models.IdRequest) (string, error)
}

type StaffI interface {
	CreateStaff(models.CreateStaffRequest) (string, error)
	UpdateStaff(models.Staff) (string, error)
	GetStaff(models.GetStaffByIdRequest) (models.Staff, error)
	GetAllStaff(models.GetAllStaffRequest) (models.GetAllStaffResponse, error)
	DeleteStaff(models.GetStaffByIdRequest) (string, error)
}

type SalesI interface {
	CreateSale(models.CreateSales) (string, error)
	UpdateSale(models.Sales) (string, error)
	GetSale(models.IdRequest) (models.Sales, error)
	GetAllSale(models.GetAllSalesRequest) (models.GetAllSalesResponse, error)
	DeleteSale(models.IdRequest) (string, error)
}
type TransactionI interface {
	CreateTransaction(models.CreateTransaction) (string, error)
	UpdateTransaction(models.Transaction) (string, error)
	GetTransaction(models.IdRequest) (models.Transaction, error)
	GetAllTransaction(models.GetAllTransactionRequest) (models.GetAllTransactionResponse, error)
	DeleteTransaction(models.IdRequest) (string, error)
}
