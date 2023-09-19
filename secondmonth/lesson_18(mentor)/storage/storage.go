package storage

import "projectWithGinAndSwagger/models"

type StorageI interface {
	Branch() BranchesI
	Staff() StaffesI
	Sales() SalesI
	Transaction() TransactionI
	StaffTarif() StaffTarifsI
	Close()
}

type BranchesI interface {
	CreateBranch(*models.CreateBranch) (string, error)
	GetBranch(*models.IdRequest) (*models.Branch, error)
	GetAllBranch(*models.GetAllBranchRequest) (*models.GetAllBranch, error)
	UpdateBranch(*models.Branch) (string, error)
	DeleteBranch(*models.IdRequest) (string, error)
}

type StaffesI interface {
	CreateStaff(*models.CreateStaff) (string, error)
	UpdateStaff(*models.Staff) (string, error)
	GetStaff(*models.IdRequest) (*models.Staff, error)
	// GetByLogin(models.LoginRequest) (*models.Staff, error)
	GetAllStaff(*models.GetAllStaffRequest) (*models.GetAllStaff, error)
	DeleteStaff(*models.IdRequest) (string, error)
	// ChangeBalance(*models.ChangeBalance) (string, error)
	// Exists(models.ExistsReq) bool
}

type TransactionI interface {
	CreateTransaction(*models.CreateTransaction) (string, error)
	UpdateTransaction(*models.Transaction) (string, error)
	GetTransaction(*models.IdRequest) (*models.Transaction, error)
	GetAllTransaction(*models.GetAllTransactionRequest) (resp *models.GetAllTransactionResponse, err error)
	DeleteTransaction(*models.IdRequest) (string, error)
	GetTopStaffs(models.TopWorkerRequest) (map[string]models.StaffTop, error)
}

type SalesI interface {
	CreateSale(*models.CreateSale) (string, error)
	UpdateSale(*models.Sales) (string, error)
	GetSale(*models.IdRequest) (*models.Sales, error)
	GetAllSale(*models.GetAllSaleRequest) (*models.GetAllSaleResponse, error)
	DeleteSale(*models.IdRequest) (string, error)
	// GetTopSaleBranch() (resp map[string]models.SaleTopBranch, err error)
	// GetSaleCountBranch() (resp map[string]models.SaleCountSumBranch, err error)
	// CancelSale(req *models.IdRequest) (string, error)
}

type StaffTarifsI interface {
	CreateStaffTarif(*models.CreateStaffTarif) (string, error)
	UpdateStaffTarif(*models.StaffTarif) (string, error)
	GetStaffTarif(*models.IdRequest) (resp *models.StaffTarif, err error)
	GetAllStaffTarif(*models.GetAllStaffTarifRequest) (resp *models.GetAllStaffTarif, err error)
	DeleteStaffTarif(*models.IdRequest) (string, error)
}
