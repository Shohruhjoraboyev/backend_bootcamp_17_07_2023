package storage

import (
	"context"
	"projectWithGinAndSwagger/models"
)

type StorageI interface {
	Branch() BranchesI
	Staff() StaffesI
	Sales() SalesI
	Transaction() TransactionI
	StaffTarif() StaffTarifsI
	BiznesLogic() BiznesLogicI
	Close()
}

type BranchesI interface {
	CreateBranch(context.Context, *models.CreateBranch) (string, error)
	GetBranch(context.Context, *models.IdRequest) (*models.Branch, error)
	GetAllBranch(context.Context, *models.GetAllBranchRequest) (*models.GetAllBranch, error)
	UpdateBranch(context.Context, *models.Branch) (string, error)
	DeleteBranch(context.Context, *models.IdRequest) (string, error)
}

type StaffesI interface {
	CreateStaff(*models.CreateStaff) (string, error)
	UpdateStaff(*models.Staff) (string, error)
	GetStaff(*models.IdRequest) (*models.Staff, error)
	GetByLogin(*models.LoginRequest) (*models.Staff, error)
	GetAllStaff(*models.GetAllStaffRequest) (*models.GetAllStaff, error)
	DeleteStaff(*models.IdRequest) (string, error)
	// ChangeBalance(*models.ChangeBalance) (string, error)
	// Exists(models.ExistsReq) bool
}

type TransactionI interface {
	CreateTransaction(context.Context, *models.CreateTransaction) (string, error)
	UpdateTransaction(context.Context, *models.Transaction) (string, error)
	GetTransaction(context.Context, *models.IdRequest) (*models.Transaction, error)
	GetAllTransaction(context.Context, *models.GetAllTransactionRequest) (resp *models.GetAllTransactionResponse, err error)
	DeleteTransaction(context.Context, *models.IdRequest) (string, error)
	// GetTopStaffs(context.Context, models.TopWorkerRequest) (map[string]models.StaffTop, error)
}

type SalesI interface {
	CreateSale(context.Context, *models.CreateSale) (string, error)
	UpdateSale(context.Context, *models.Sales) (string, error)
	GetSale(context.Context, *models.IdRequest) (*models.Sales, error)
	GetAllSale(context.Context, *models.GetAllSaleRequest) (*models.GetAllSaleResponse, error)
	DeleteSale(context.Context, *models.IdRequest) (string, error)
	// GetTopSaleBranch() (resp map[string]models.SaleTopBranch, err error)
	// GetSaleCountBranch() (resp map[string]models.SaleCountSumBranch, err error)
	// CancelSale(req *models.IdRequest) (string, error)
}

type StaffTarifsI interface {
	CreateStaffTarif(context.Context, *models.CreateStaffTarif) (string, error)
	UpdateStaffTarif(context.Context, *models.StaffTarif) (string, error)
	GetStaffTarif(context.Context, *models.IdRequest) (resp *models.StaffTarif, err error)
	GetAllStaffTarif(context.Context, *models.GetAllStaffTarifRequest) (resp *models.GetAllStaffTarif, err error)
	DeleteStaffTarif(context.Context, *models.IdRequest) (string, error)
}

type BiznesLogicI interface {
	GetTopStaffs(*models.TopStaffRequest) (resp *models.TopStaffResponse, err error)
}
