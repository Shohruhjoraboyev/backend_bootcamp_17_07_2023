package storage

import (
	"app/models"
)

type StorageI interface {
	Branch() BranchesI
	Tariff() TariffsI
	Staff() StaffesI
	// Sales() SalesI
	// Transaction() TransactionI
	// StaffTarif() StaffTarifsI
	Close()
}

type BranchesI interface {
	CreateBranch(*models.CreateBranch) (string, error)
	GetBranch(*models.IdRequest) (*models.Branch, error)
	GetAllBranch(*models.GetAllBranchRequest) (*models.GetAllBranch, error)
	UpdateBranch(*models.Branch) (string, error)
	DeleteBranch(*models.IdRequest) (string, error)
}

type TariffsI interface {
	CreateStaffTarif(*models.CreateStaffTarif) (string, error)
	GetStaffTarif(*models.IdRequest) (*models.StaffTarif, error)
	GetAllStaffTarif(*models.GetAllStaffTarifRequest) (*models.GetAllStaffTarif, error)
	UpdateStaffTarif(*models.StaffTarif) (string, error)
	DeleteStaffTarif(*models.IdRequest) (string, error)
}

type StaffesI interface {
	CreateStaff(*models.CreateStaff) (string, error)
	UpdateStaff(*models.Staff) (string, error)
	GetStaff(*models.IdRequest) (*models.Staff, error)
	// GetByLogin(models.LoginRequest) (models.Staff, error)
	GetAllStaff(*models.GetAllStaffRequest) (*models.GetAllStaff, error)
	DeleteStaff(*models.IdRequest) (string, error)
	// ChangeBalance(models.ChangeBalance) (string, error)
	// Exists(models.ExistsReq) bool
}
